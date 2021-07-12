package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLDao struct {
	conn *gorm.DB
}

func newSQLDao() (error, SQLDao) {
	connection, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err, SQLDao{}
	}

	return err, SQLDao{connection.LogMode(true)}
}

func (dao SQLDao) Close() {
	dao.conn.Close()
}

func WithTx(runner func(dao SQLDao) error) error {
	conn, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	runner(SQLDao{tx.LogMode(true)})

	return tx.Commit().Error
}
