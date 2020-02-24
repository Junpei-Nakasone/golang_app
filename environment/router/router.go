package router

import (
	"go_app/api/api001"

	"github.com/labstack/echo"
)

// NewRouter return router
func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", api001.Test)

	return e
}
