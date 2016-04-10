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
			"error": err,
		}).Error("error parsing request")
		return
	}
	email, ok := user_info["email"]
	if !ok {
		email_error := "no email provided"
		c.JSON(500, gin.H{"error": email_error})
		log.WithFields(log.Fields{
			"error": email_error,
		}).Info("error creating user")
	} else {
		err := database.CreateUser(email)
		if err == nil {
			c.JSON(200, gin.H{"success": "true"})
			log.WithFields(log.Fields{
				"email": email,
			}).Info("created user")
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Info("error creating user")
		}
	}
}
