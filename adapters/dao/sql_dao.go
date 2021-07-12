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
	conn *TxSQLDao
}

func WITHOUT_TX() txType {
	return txType{}
}

func WITH_TX(tx TxSQLDao) txType {
	return txType{&tx}
}

func newSQLDao(tt txType) (error, SQLDao) {
	if tt.conn != nil {
		return nil, tt.conn.value
	}

	connection, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err, SQLDao{}
	}

	return err, SQLDao{connection.LogMode(true)}
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
		return tx.Error
	}

	runner(TxSQLDao{SQLDao{tx.LogMode(true)}})

	return tx.Commit().Error
}
