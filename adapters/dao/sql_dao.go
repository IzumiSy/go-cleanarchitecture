package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	Dialector() gorm.Dialector
}

func newSQLDao(tableName string, tt txType) (SQLDao, error) {
	if tt.dao != nil {
		tt.dao.value.conn.Logger = logger.Default
		return SQLDao{tt.dao.value.conn.Table(tableName)}, nil
	}

	connection, err := gorm.Open(currentDriver().Dialector(), &gorm.Config{})
	if err != nil {
		return SQLDao{}, err
	}

	connection.Logger = logger.Default
	return SQLDao{connection.Table(tableName)}, nil
}

func WithTx(runner func(tx TxSQLDao) error) error {
	conn, err := gorm.Open(currentDriver().Dialector(), &gorm.Config{})
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
