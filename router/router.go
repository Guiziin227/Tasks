package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	log.Print("Server is running on port 8080")
	r.Run(":8080")
}
