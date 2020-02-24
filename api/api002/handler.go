package api002

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

// NewUser add new user
func NewUser(c echo.Context) error {
	db := db.CreateDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// id := c.QueryParam("id")
	// name := c.QueryParam("name")
	// age := c.QueryParam("age")

	var users Users

	c.Bind(&users)

	// var users []Users
	db.Create(&Users{ID: users.ID, Name: users.Name, Age: users.Age})

	return c.JSON(http.StatusCreated, "new user added")
}
