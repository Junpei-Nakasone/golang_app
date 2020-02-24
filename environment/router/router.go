package router

import (
	"go_app/api/api001"
	"go_app/api/api002"
	"go_app/api/api003"

	"github.com/labstack/echo"
)

// NewRouter return router
func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", api001.Test)
	e.POST("/newUser", api002.NewUser)
	e.POST("/updateUser", api003.UpdateUser)

	return e
}
