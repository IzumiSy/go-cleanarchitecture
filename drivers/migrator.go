package drivers

import (
	"context"
	"fmt"
	"go-cleanarchitecture/adapters/dao"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MigratorDriver struct {
	Mode string
}

func (driver MigratorDriver) Run(ctx context.Context) {
	conn, err := gorm.Open(dao.CurrentDriver().Dialector(), &gorm.Config{})
	conn.Logger = logger.Default.LogMode(logger.Info)
	if err != nil {
		panic(err.Error())
	}

	switch driver.Mode {
	case "down":
		driver.Down(conn)
	case "up":
		driver.Up(conn)
	default:
		fmt.Println("Migration supports only `up` or `down`")
		os.Exit(1)
	}
}

func (driver MigratorDriver) Up(conn *gorm.DB) {
	conn.Migrator().CreateTable(&dao.TodoDto{})
	conn.Migrator().CreateTable(&dao.AuthenticationDto{})
	conn.Migrator().CreateTable(&dao.SessionDto{})
	conn.Migrator().CreateTable(&dao.UserDto{})
}

func (driver MigratorDriver) Down(conn *gorm.DB) {
	conn.Migrator().DropTable(&dao.TodoDto{})
	conn.Migrator().DropTable(&dao.AuthenticationDto{})
	conn.Migrator().DropTable(&dao.SessionDto{})
	conn.Migrator().DropTable(&dao.UserDto{})
}
