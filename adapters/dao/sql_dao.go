package dao

import (
	"log"
	"os"
	"time"

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

var appLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	},
)

func newSQLDao(tableName string, tt txType) (SQLDao, error) {
	if tt.dao != nil {
		tt.dao.value.conn.Logger = appLogger
		return SQLDao{tt.dao.value.conn.Table(tableName)}, nil
	}

	connection, err := gorm.Open(CurrentDriver().Dialector(), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return SQLDao{}, err
	}

	connection.Logger = appLogger
	return SQLDao{connection.Table(tableName)}, nil
}

func WithTx(runner func(tx TxSQLDao) error) error {
	conn, err := gorm.Open(CurrentDriver().Dialector(), &gorm.Config{
		SkipDefaultTransaction: true,
	})
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
