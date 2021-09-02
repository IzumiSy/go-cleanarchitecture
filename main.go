package main

import (
	"context"
	"flag"
	"fmt"
	adapters "go-cleanarchitecture/adapters"
	"os"
)

type Driver interface {
	Run(ctx context.Context)
}

func main() {
	migrationModePtr := flag.Bool("migrate", false, "migration mode")
	httpPtr := flag.Bool("http", false, "http server mode")
	flag.Parse()

	var driver Driver
	if migrationModePtr != nil && *migrationModePtr {
		driver = adapters.MigratorDriver{}
	} else if httpPtr != nil && *httpPtr {
		driver = adapters.HttpDriver{}
	} else {
		fmt.Println("Error: Unsupported mode specified")
		os.Exit(1)
	}

	ctx := context.Background() // TODO: timeout required
	driver.Run(ctx)
}
