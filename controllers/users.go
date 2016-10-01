package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
)

func createUser(c *gin.Context) {
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}

	username, ok := user_info["username"]
	if !ok {
		username_error := "no username provided"
		c.JSON(500, gin.H{"error": username_error})
		log.WithFields(log.Fields{
			"at":    "controllers.CreateUser",
			"error": username_error,
		}).Error("error creating user")
		return
	}

	reset_link, err := models.CreateUser(username)
	if err == nil {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at":         "controllers.CreateUser",
			"username":   username,
			"reset_link": reset_link,
		}).Info("created user")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.CreateUser",
			"username": username,
			"error":    err.Error(),
		}).Error("error creating user")
	}
}

func updateUserName(c *gin.Context) {
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}

	name, ok := user_info["name"]
	if !ok {
		name_error := "no name provided"
		c.JSON(500, gin.H{"error": name_error})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserName",
			"error": name_error,
		}).Error("error updating user's name")
		return
	}

	user, err := currentUser(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserName",
			"error": err.Error(),
		}).Error("error getting current user")
		return
	}

	user.Name = name
	db_err := user.Save()
	if db_err != nil {
		c.JSON(500, gin.H{"error": db_err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserName",
			"error": db_err.Error(),
		}).Error("error updating user")
		return
	} else {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at": "controllers.UpdateUserName",
		}).Info("user name updated")
	}
}

func passwordReset(c *gin.Context) {

}

func updateUserPassword(c *gin.Context) {
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}
	user, err := currentUser(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	old_password, ok := user_info["old_password"]
	if !ok {
		// no old password provided
	}

	// ensure old password correct password
	if !user.ValidAuthentication(old_password) {

	}

	new_password, ok := user_info["new_password"]
	if !ok {
		name_error := "no name provided"
		c.JSON(500, gin.H{"error": name_error})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserName",
			"error": name_error,
		}).Error("error updating user's name")
		return
	}

	// Set new password
	err = user.SetPassword(new_password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserPassword",
			"error": err.Error(),
		}).Error("err setting user password")
		return
	} else {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at": "controllers.UpdateUserName",
		}).Info("user password updated")
	}
}

func destroyUser(c *gin.Context) {

}

//
//
//
//

func currentUser(c *gin.Context) (user models.User, err error) {
	session_cookie, err := sessionCookie(c)
	if err != nil {
		return
	}
	user, err = userFromSessionCookie(session_cookie)
	return
}

func userFromSessionCookie(session_cookie string) (user models.User, err error) {
	session, err := models.SessionFromID(session_cookie)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "database.currentUser",
			"error": err,
		}).Error("session_missing")
		return
	}
	user, err = session.User()
	//
	return
}

func requestJSON(c *gin.Context) (json map[string]string, err error) {
	err = c.BindJSON(&json)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		log.WithFields(log.Fields{
			"at":    "controllers.requestJSON",
			"error": err,
		}).Error("error parsing request")
	}
	return
}
