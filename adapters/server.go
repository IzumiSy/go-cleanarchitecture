package adapters

import (
	"context"
	"go-cleanarchitecture/adapters/loggers"
	"go-cleanarchitecture/adapters/pubsub"
	"go-cleanarchitecture/domains"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunHTTPServer(ctx context.Context) {
	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		panic(err)
	}

	err, pa := pubsub.NewRedisAdapter(logger)
	if err != nil {
		panic(err)
	}

	pa.RegisterSubscriber(domains.UserSignedUp, signedUpHandler(ctx, logger))
	pa.RegisterSubscriber(domains.UserAuthenticated, userAuthenticatedHandler(ctx, logger))
	pa.RegisterSubscriber(domains.TodoCreated, todoCreatedHandler(ctx, logger))
	go pa.Listen(ctx)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))

	e.GET("/todos", getTodosHandler(logger))
	e.POST("/todo", createTodoHandler(pa, logger))
	e.POST("/signup", signupHandler(pa, logger))
	e.POST("/login", authenticateHandler(pa, logger))

	e.Logger.Fatal(e.Start(":8080"))
}
