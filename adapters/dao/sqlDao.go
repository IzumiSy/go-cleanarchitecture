package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func WithConnection(f func(conn *gorm.DB) error) error {
	connection, err := gorm.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		return err
	}
	defer connection.Close()

	return f(connection)
}
