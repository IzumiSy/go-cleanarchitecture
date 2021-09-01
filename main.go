package main

import (
	"context"
	"flag"
	"fmt"
	drivers "go-cleanarchitecture/drivers"
	"os"
)

type Driver interface {
	Run(ctx context.Context)
}

func main() {
	migrationModePtr := flag.String("migrate", "", "migration mode")
	httpPtr := flag.Bool("http", false, "http server mode")
	flag.Parse()

	var driver Driver
	if migrationModePtr != nil && *migrationModePtr != "" {
		driver = drivers.MigratorDriver{Mode: *migrationModePtr}
	} else if httpPtr != nil && *httpPtr {
		driver = drivers.HttpDriver{}
	} else {
		fmt.Println("Error: Unsupported mode specified")
		os.Exit(1)
	}

	ctx := context.Background() // TODO: timeout required
	driver.Run(ctx)
}
