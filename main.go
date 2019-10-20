package main

import (
	"database/sql"
	"flag"
	"fmt"
	"go-cleanarchitecture/adapters"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	migratorPtr := flag.Bool("migrate", false, "migration mode")
	flag.Parse()

	if migratorPtr != nil && *migratorPtr {
		migration()
		return
	}

	adapters.RunHTTPServer()
}

func migration() {
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
