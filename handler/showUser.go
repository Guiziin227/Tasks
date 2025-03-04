package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tasks/db"
	"tasks/models"
)

func ShowUser(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// Find user
	user := models.User{}
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
