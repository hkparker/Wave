package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
	"net/http"
)

// Handle login requests
func newSession(c *gin.Context) {
	// Ensure the request is valid JSON
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}

	// Ensure the JSON contains a username key with data
	username, ok := user_info["username"]
	if !ok || username == "" {
		err := "no user provided"
		c.JSON(400, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.newSession",
			"error": err,
		}).Error("error creating session")
		return
	}

	// Ensure the JSON contains a password key
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

	// Ensure the user exists
	user, err := models.UserByUsername(username)
	if err != nil {
		err := "user does not exist"
		// Do not reveal if the username was a valid account
		c.JSON(401, gin.H{"error": "incorrect authentication"})
		log.WithFields(log.Fields{
			"at":       "controllers.newSession",
			"error":    err,
			"username": username,
		}).Error("error creating session")
		return

	}

	// Check if the provided password is correct
	if user.ValidAuthentication(password) {
		// Create a new sessions
		session_cookie, err := user.NewSession()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			log.WithFields(log.Fields{
				"at":       "controllers.newSession",
				"error":    err.Error(),
				"username": username,
			}).Error("error creating session")
			return
		}
		session, err := models.SessionFromID(session_cookie)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			log.WithFields(log.Fields{
				"at":       "controllers.newSession",
				"error":    err.Error(),
				"username": username,
			}).Error("error looking up new sessions")
			return
		}

		// Set the wave_session cookie
		cookie := session.HTTPCookie()
		http.SetCookie(c.Writer, &cookie)
		c.JSON(200, gin.H{"authentication": "success"})
		log.WithFields(log.Fields{
			"at":       "controllers.newSession",
			"username": username,
		}).Info("session created")
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

func deleteSession(c *gin.Context) {
	session_cookie, serr := sessionCookie(c)

	if serr == nil {
		session, err := models.SessionFromID(session_cookie)
		if err != nil {
			c.JSON(401, gin.H{"error": "no session"})
			log.WithFields(log.Fields{
				"at":    "controllers.deleteSession",
				"error": err.Error(),
			}).Error("error finding session for cookie")
			return
		}
		err = session.Delete()
		if err != nil {
			c.JSON(500, gin.H{"error": "error deleting session"})
			log.WithFields(log.Fields{
				"at":    "controllers.deleteSession",
				"error": err.Error(),
			}).Error("error deleting session")
			return
		}
	}

	c.Redirect(302, "/login")
	c.Abort()
}
