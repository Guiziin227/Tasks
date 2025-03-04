package main

import (
	"fmt"
	"tasks/db"
	"tasks/router"
)

func main() {

	err := db.Connect()
	if err != nil {
		fmt.Sprint("Error initializing config: %v", err)
		return
	}

	router.Initialize()

}
