package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rubenv/sql-migrate"
)

func main() {
	migrations := &migrate.FileMigrationSource{Dir: "schemas"}

	db, err := sql.Open("sqlite3", "go-cleanarchitecture.db")
	if err != nil {
		panic(err.Error())
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Applied %d migrations\n", n)
}
