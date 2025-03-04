package router

import (
	"github.com/gin-gonic/gin"
	"tasks/handler"
)

func initializeRoutes(router *gin.Engine) {

	user := router.Group("/api/")
	{
		//Show opening
		user.GET("/users", handler.ShowUsers)
		user.GET("/users/", handler.ShowUser)
		user.POST("/users", handler.CreateUser)
		user.DELETE("/users/", handler.DeleteUser)
	}

}
