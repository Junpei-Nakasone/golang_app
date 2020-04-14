package api001

import (
	"net/http"

	"github.com/todolist_ver2/environment/db"

	"github.com/labstack/echo"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Users 構造体
type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var err error

// Test returns json
func Test(c echo.Context) error {
	db := db.CreateDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	var users []Users
	db.Find(&users)
	return c.JSON(http.StatusOK, &users)
}
