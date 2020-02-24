package main

import (
	"go_app/environment/router"
)

func main() {
	e := router.NewRouter()

	e.Start(":8008")
}
