package api004

import (
	"go_app/environment/db"
	"net/http"

	"github.com/labstack/echo"
)

// Users 構造体
type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var err error


// DeleteUser delete user
func DeleteUser(c echo.Context) error {
	db := db.CreateDBConnection()
	if err != nil {
		panic("could not connect to DB.")
	}
	defer db.Close()

	var user Users

	c.Bind(&user)

	db.Where("name = ?", user.Name).Find(&user)
	db.Delete(&user)

	return c.JSON(http.StatusNoContent, "user deleted")
}