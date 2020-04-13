package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLDao struct {
	conn *gorm.DB
}

func newSQLDao(table string) (error, SQLDao) {
	connection, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err, SQLDao{}
	}

	return err, SQLDao{connection.LogMode(true).Table(table)}
}

func (dao SQLDao) Close() {
	dao.conn.Close()
}
