package adapters

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/loggers"
	"go-cleanarchitecture/adapters/pubsub"
	"go-cleanarchitecture/domains"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

type HttpDriver struct{}

func (driver HttpDriver) Run(ctx context.Context) {
	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		panic(err)
	}

	err, pa := pubsub.NewRedisAdapter(logger)
	if err != nil {
		panic(err)
	}

	pa.RegisterSubscriber(domains.UserSignedUp, signedUpSubscriber(ctx, logger))
	pa.RegisterSubscriber(domains.UserAuthenticated, userAuthenticatedSubscriber(ctx, logger))
	pa.RegisterSubscriber(domains.TodoCreated, todoCreatedSubscriber(ctx, logger))
	go pa.Listen(ctx)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))

	e.GET("/todos", getTodosController(logger))
	e.POST("/todo", createTodoController(pa, logger))
	e.POST("/signup", signupController(pa, logger))
	e.POST("/login", authenticateController(pa, logger))

	e.Logger.Fatal(e.Start(":8080"))
}

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
