package main

import (
	"go_app/environment/router"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mysql")
)

func main() {
	// db := db.CreateDBConnection()
	// fmt.Println(db)
	e := router.NewRouter()

	e.Start(":8000")
}
