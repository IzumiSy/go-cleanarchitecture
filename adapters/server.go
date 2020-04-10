package adapters

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RunHTTPServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todos", getTodosHandler)
	e.POST("/todo", createTodoHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
