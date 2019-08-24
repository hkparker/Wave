package controllers

import (
	//log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
)

func getStatus(c *gin.Context) {
	userInterface, ok := c.Get("currentUser")
	if !ok {
		c.Status(500)
		return
	}
	user, ok := userInterface.(models.User)
	if !ok {
		c.Status(500)
		return
	}
	
	c.JSON(200, gin.H{"user": user.Username})
}
