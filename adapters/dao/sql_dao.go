package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLDao struct {
	conn *gorm.DB
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

type driverLike interface {
	Dialect() string
	DSN() string
}

func newSQLDao(tableName string, tt txType, driver driverLike) (SQLDao, error) {
	if tt.dao != nil {
		return SQLDao{tt.dao.value.conn.LogMode(true).Table(tableName)}, nil
	}

	connection, err := gorm.Open(driver.Dialect(), driver.DSN())
	if err != nil {
		return SQLDao{}, err
	}

	return SQLDao{connection.LogMode(true).Table(tableName)}, nil
}

func (dao SQLDao) Close() {
	dao.conn.Close()
}

func WithTx(runner func(tx TxSQLDao) error) error {
	conn, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err
	}

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
