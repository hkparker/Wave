package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
	"net/http"
)

func newSession(c *gin.Context) {
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}

	username, ok := user_info["username"]
	if !ok {
		err := "no user provided"
		c.JSON(400, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.newSession",
			"error": err,
		}).Error("error creating session")
		return
	}

	password, ok := user_info["password"]
	if !ok {
		err := "no password provided"
		c.JSON(400, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":       "controllers.newSession",
			"error":    err,
			"username": username,
		}).Error("error creating session")
		return
	}

	user, err := models.UserByUsername(username)
	if err != nil {
		err := "user does not exist"
		c.JSON(401, gin.H{"error": "incorrect authentication"})
		log.WithFields(log.Fields{
			"at":       "controllers.newSession",
			"error":    err,
			"username": username,
		}).Error("error creating session")
		return

	}

	if user.ValidAuthentication(password) {
		session_cookie, err := user.NewSession()
		if err != nil {

		}
		session, err := models.SessionFromID(session_cookie)
		if err != nil {

		}

		cookie := session.HTTPCookie()
		http.SetCookie(c.Writer, &cookie)
		c.JSON(200, gin.H{"authentication": "success"})
		log.WithFields(log.Fields{
			"at":       "controllers.newSession",
			"username": username,
		}).Error("session created")
	} else {
		err := "incorrect password"
		c.JSON(401, gin.H{"error": "incorrect authentication"})
		log.WithFields(log.Fields{
			"at":       "controllers.newSession",
			"error":    err,
			"username": username,
		}).Error("error creating session")
	}
}
