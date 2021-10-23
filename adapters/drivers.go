package adapters

import (
	"context"
	"fmt"
	"go-cleanarchitecture/adapters/dao"
	"go-cleanarchitecture/adapters/loggers"
	"go-cleanarchitecture/adapters/pubsub"
	"go-cleanarchitecture/domains"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type driverConfigs struct {
	Name  string
	DB    dao.Driver
	Redis string
}

// default: dev
var currentDriver driverConfigs = driverConfigs{
	Name:  "development",
	DB:    dao.DevDriver,
	Redis: "localhost:6379",
}

func loadConfigs() {
	switch os.Getenv("APP_ENV") {
	case "production":
		currentDriver = driverConfigs{
			Name:  "production",
			DB:    dao.ProdDriver,
			Redis: "redis:6379",
		}
	}
}

type HttpDriver struct{}

func (driver HttpDriver) Run(ctx context.Context) {
	loadConfigs()

	logger, err := loggers.NewZapLogger()
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "Env: %s", currentDriver.Name)

	err, pa := pubsub.NewRedisAdapter(logger, currentDriver.Redis)
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

	v1 := e.Group("v1")
	v1.GET("/todos", getTodosController(logger, currentDriver.DB))
	v1.POST("/todo", createTodoController(pa, logger, currentDriver.DB))
	v1.POST("/signup", signupController(pa, logger, currentDriver.DB))
	v1.POST("/login", authenticateController(pa, logger, currentDriver.DB))

	e.Logger.Fatal(e.Start(":8080"))
}

type MigratorDriver struct{}

func (driver MigratorDriver) Run(ctx context.Context) {
	fmt.Println("Deprecated driver")
}
