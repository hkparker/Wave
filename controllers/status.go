package controllers

import (
	//log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func getStatus(c *gin.Context) {
	user, err := currentUser(c)
	if err != nil {
		return
	}
	
	c.JSON(200, gin.H{
		"username": user.Username,
		"admin": user.Admin,
		"version": 0,
	})
}
