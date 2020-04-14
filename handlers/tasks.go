package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/todolist_ver2/models"
)

// GetTasks endpoint
func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTodos())
}
