package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Dao

type SQLDao struct {
	conn *gorm.DB
}

type Driver struct {
	Dialector gorm.Dialector
}

var (
	DevDriver  = Driver{Dialector: sqlite.Open("go-cleanarchitecture.db")}
	ProdDriver = Driver{Dialector: mysql.Open("root:password@tcp(db:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local")}
)

var defaultLogger = logger.Default.LogMode(logger.Info)

func (driver Driver) newSQLDao(tableName string, tt txType) (SQLDao, error) {
	if tt.dao != nil {
		return SQLDao{tt.dao.value.conn.Table(tableName)}, nil
	}

	connection, err := gorm.Open(driver.Dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return SQLDao{}, err
	}

	connection.Logger = defaultLogger
	return SQLDao{connection.Table(tableName)}, nil
}

// Transaction

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

func (driver Driver) WithTx(runner func(tx TxSQLDao) error) error {
	conn, err := gorm.Open(driver.Dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}

	conn.Logger = defaultLogger
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
