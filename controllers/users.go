package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/database"
)

func CreateUser(c *gin.Context) {
	var user_info map[string]string
	err := c.BindJSON(&user_info)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.CreateUser",
			"error": err,
		}).Error("error parsing request")
		return
	}

	email, ok := user_info["email"]
	if !ok {
		email_error := "no email provided"
		c.JSON(500, gin.H{"error": email_error})
		log.WithFields(log.Fields{
			"at":    "controllers.CreateUser",
			"error": email_error,
		}).Error("error creating user")
		return
	}

	err = database.CreateUser(email)
	if err == nil {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at":    "controllers.CreateUser",
			"email": email,
		}).Info("created user")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.CreateUser",
			"email": email,
			"error": err.Error(),
		}).Error("error creating user")
	}
}
