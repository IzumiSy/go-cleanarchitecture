package main

import (
	"flag"
	"fmt"
	drivers "go-cleanarchitecture/drivers"
	"os"
)

type Driver interface {
	Run()
}

func main() {
	migrationModePtr := flag.String("migrate", "up", "migration mode")
	httpPtr := flag.Bool("http", false, "http server mode")
	flag.Parse()

	var driver Driver
	if migrationModePtr != nil {
		driver = drivers.MigratorDriver{Mode: *migrationModePtr}
	} else if httpPtr != nil && *httpPtr {
		driver = drivers.HttpDriver{}
	} else {
		fmt.Println("Error: Unsupported mode specified")
		os.Exit(1)
	}

	driver.Run()
}
