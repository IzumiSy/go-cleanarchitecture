package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLDao struct {
	conn *gorm.DB
}

func (dao SQLDao) Table(name string) {
	dao.conn.Table(name)
}

type TxSQLDao struct {
	value SQLDao
}

type txType struct {
	dao *TxSQLDao
}

func WITHOUT_TX() txType {
	return txType{}
}

func WITH_TX(tx TxSQLDao) txType {
	return txType{&tx}
}

func newSQLDao(tableName string, tt txType) (error, SQLDao) {
	if tt.dao != nil {
		return nil, SQLDao{tt.dao.value.conn.LogMode(true).Table(tableName)}
	}

	connection, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err, SQLDao{}
	}

	return err, SQLDao{connection.LogMode(true).Table(tableName)}
}

func (dao SQLDao) Close() {
	dao.conn.Close()
}

func WithTx(runner func(tx TxSQLDao) error) error {
	conn, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err
	}
	conn.LogMode(true)

	tx := conn.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	if err := runner(TxSQLDao{SQLDao{tx}}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
