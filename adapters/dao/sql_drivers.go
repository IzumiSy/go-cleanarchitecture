package dao

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type driver struct {
	dialector gorm.Dialector
}

func (d driver) Dialector() gorm.Dialector {
	return d.dialector
}

func CurrentDriver() driver {
	switch os.Getenv("APP_ENV") {
	case "production":
		return ProdDriver
	default:
		return DevDriver
	}
}

var DevDriver = driver{
	dialector: sqlite.Open("go-cleanarchitecture.db"),
}

var ProdDriver = driver{
	dialector: mysql.Open("root:password@tcp(db:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"),
}
