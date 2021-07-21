package dao

import "os"

type driver struct {
	dialect string
	dsn     string
}

func (d driver) Dialect() string {
	return d.dialect
}

func (d driver) DSN() string {
	return d.dsn
}

func currentDriver() driver {
	switch os.Getenv("APP_ENV") {
	case "production":
		return ProdDriver
	default:
		return DevDriver
	}
}

var DevDriver = driver{
	dialect: "sqlite3",
	dsn:     "go-cleanarchitecture.db",
}

var ProdDriver = driver{
	dialect: "mysql",
	dsn:     "root:password@tcp(localhost:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local",
}
