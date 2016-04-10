package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/database"
)

func CreateUser(c *gin.Context) {
	//c.Params()
	err := database.CreateUser("fixit")
	if err == nil {
		c.JSON(200, gin.H{"success": "true"})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}
