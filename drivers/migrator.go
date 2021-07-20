package drivers

import (
	"go-cleanarchitecture/adapters/dao"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MigratorDriver struct {
	Mode string
}

func (driver MigratorDriver) Run() {
	conn, err := gorm.Open(dao.CurrentDriver().Dialector(), &gorm.Config{})
	conn.Logger = logger.Default.LogMode(logger.Info)
	if err != nil {
		panic(err.Error())
	}

	switch driver.Mode {
	case "down":
		driver.Down(conn)
	default:
		driver.Up(conn)
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
