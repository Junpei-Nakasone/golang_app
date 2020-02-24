package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var connection *gorm.DB
var once sync.Once

// NewDBConnection connect DB
// func NewDBConnection() *gorm.DB {
// 	once.Do(func(){
// 		connection = createDBConnection()
// 	})
// 	return connection
// }

// CreateDBConnection create conn
func CreateDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mysql")
	// if db.DB().Ping() != nil{
	// 	panic(err)
	// }
	if err != nil {
		return db
	}
	return db
}