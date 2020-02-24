package api001

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// TaskHandler gets Db
type TaskHandler struct {
	Db *gorm.DB
}

// Users 構造体
type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	age  int    `json:"age"`
}

// Test returns json
func (handler *TaskHandler) Test(c echo.Context) error {
	handler.Db.Find(&Users)
	return c.JSON(http.StatusOK, &Users)
}
