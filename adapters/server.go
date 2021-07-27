package adapters

import (
	"go-cleanarchitecture/adapters/pubsub"
	"go-cleanarchitecture/domains"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RunHTTPServer() {
	err, pa := pubsub.NewRedisAdapter()
	if err != nil {
		panic(err)
	}

	pa.RegisterSubscriber(domains.UserSignedUp, signedUpHandler)
	pa.RegisterSubscriber(domains.UserAuthenticated, userAuthenticatedHandler)
	pa.RegisterSubscriber(domains.TodoCreated, todoCreatedHandler)
	go pa.Listen()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todos", getTodosHandler)
	e.POST("/todo", createTodoHandler(pa))
	e.POST("/signup", signupHandler(pa))
	e.POST("/login", authenticateHandler(pa))

	e.Logger.Fatal(e.Start(":8080"))
}
