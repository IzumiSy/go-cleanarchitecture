package main

import (
	"flag"
	"fmt"
	drivers "go-cleanarchitecture/drivers"
	"os"
)

type Driver interface {
	Run(options drivers.Options)
}

func main() {
	migratorPtr := flag.Bool("migrate", false, "migration mode")
	httpPtr := flag.Bool("http", false, "http server mode")
	flag.Parse()

	var driver Driver
	if migratorPtr != nil && *migratorPtr {
		driver = drivers.MigratorDriver{}
	} else if httpPtr != nil && *httpPtr {
		driver = drivers.HttpDriver{}
	} else {
		fmt.Println("Error: Unsupported mode specified")
		os.Exit(1)
	}

	options := drivers.Options{}
	driver.Run(options)
}
