package api002

import (
	"net/http"
	"github.com/todolist_ver2/environment/db"

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

	var users Users

	c.Bind(&users)

	tx := db.Begin()
	// defer func()も必要？
	if err := tx.Create(&Users{ID: users.ID, Name: users.Name, Age: users.Age}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	// db.Create(&Users{ID: users.ID, Name: users.Name, Age: users.Age})

	return c.JSON(http.StatusCreated, "new user added")
}
