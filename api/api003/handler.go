package api003

import (
	"go_app/environment/db"
	"net/http"

	"github.com/labstack/echo"
)

// Users define user
type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var err error

// UpdateUser updates user
func UpdateUser(c echo.Context) error {
	db := db.CreateDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	var users Users

	c.Bind(&users)

	db.Where("name = ?", users.Name).Find(&users)

	// return c.JSON(http.StatusOK, &users)
	db.Save(&Users{ID: users.ID, Name: users.Name, Age: users.Age})

	return c.JSON(http.StatusNoContent, "User updated")
}
