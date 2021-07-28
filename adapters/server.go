package adapters

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-cleanarchitecture/adapters/loggers"
	"go-cleanarchitecture/adapters/pubsub"
	"go-cleanarchitecture/domains"
)

func RunHTTPServer() {
	logger, err := loggers.NewZapLogger("config/zap.json")
	if err != nil {
		panic(err)
	}

	err, pa := pubsub.NewRedisAdapter()
	if err != nil {
		panic(err)
	}

	pa.RegisterSubscriber(domains.UserSignedUp, signedUpHandler)
	pa.RegisterSubscriber(domains.UserAuthenticated, userAuthenticatedHandler)
	pa.RegisterSubscriber(domains.TodoCreated, todoCreatedHandler)
	go pa.Listen(logger)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todos", getTodosHandler(logger))
	e.POST("/todo", createTodoHandler(pa, logger))
	e.POST("/signup", signupHandler(pa, logger))
	e.POST("/login", authenticateHandler(pa, logger))

	e.Logger.Fatal(e.Start(":8080"))
}
