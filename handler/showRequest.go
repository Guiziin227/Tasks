// handler/showRequest.go
package handler

import (
	"net/http"
	"tasks/db"
	"tasks/models"

	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	if db.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
