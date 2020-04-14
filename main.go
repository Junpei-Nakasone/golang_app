package main

import (
	"github.com/todolist_ver2/environment/router"
)

func main() {
	e := router.NewRouter()

	e.Start(":8008")
}
