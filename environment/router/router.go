package router

import (
	"github.com/todolist_ver2/handlers"

	"github.com/labstack/echo"
)

// NewRouter return router
func NewRouter() *echo.Echo {
	e := echo.New()
	e.File("/", "public/index.html")
	e.GET("/todos", handlers.GetTodos)
	return e
}
