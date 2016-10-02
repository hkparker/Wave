package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
)

// Handle a request to create a new user
func createUser(c *gin.Context) {
	// Ensure the request is valid JSON and get
	// the user from the current session
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}
	admin, err := currentUser(c)
	if err != nil {
		return
	}

	// Ensure the JSON contains a username key with data
	username, ok := user_info["username"]
	if !ok || username == "" {
		username_error := "no username provided"
		c.JSON(500, gin.H{"error": username_error})
		log.WithFields(log.Fields{
			"at":    "controllers.createUser",
			"error": username_error,
			"admin": admin.Username,
		}).Error("error creating user")
		return
	}

	reset_link, err := models.CreateUser(username)
	if err == nil {
		c.JSON(200, gin.H{
			"success":    "true",
			"reset_link": reset_link,
		})
		log.WithFields(log.Fields{
			"at":       "controllers.createUser",
			"username": username,
			"admin":    admin.Username,
		}).Info("created user")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.createUser",
			"username": username,
			"error":    err.Error(),
			"admin":    admin.Username,
		}).Error("error creating user")
	}
}

func updateUserName(c *gin.Context) {
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}
	user, err := currentUser(c)
	if err != nil {
		return
	}

	username, ok := user_info["username"]
	if !ok {
		name_error := "no name provided"
		c.JSON(500, gin.H{"error": name_error})
		log.WithFields(log.Fields{
			"at":       "controllers.updateUserName",
			"error":    name_error,
			"username": username,
			"user":     user.Name,
		}).Error("error updating user name")
		return
	}

	user.Name = username
	db_err := user.Save()
	if db_err != nil {
		c.JSON(500, gin.H{"error": db_err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.updateUserName",
			"error":    db_err.Error(),
			"username": username,
			"user":     user.Name,
		}).Error("error updating user")
		return
	} else {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at":       "controllers.updateUserName",
			"username": username,
			"old_name": user.Name,
		}).Info("user name updated")
	}
}

func updateUserPassword(c *gin.Context) {
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}
	user, err := currentUser(c)
	if err != nil {
		return
	}

	old_password, ok := user_info["old_password"]
	if !ok {
		err := "no old password provided"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.updateUserPassword",
			"error": err,
		}).Error("error updating user password")
		return
	}

	new_password, ok := user_info["new_password"]
	if !ok {
		err := "no new password provided"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.updateUserPassword",
			"error": err,
		}).Error("error updating user password")
		return
	}

	if !user.ValidAuthentication(old_password) {
		err := "old password incorrect"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.updateUserPassword",
			"error": err,
		}).Error("error updating user password")
		return
	}

	err = user.SetPassword(new_password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserPassword",
			"error": err.Error(),
		}).Error("err setting user password")
		return
	} else {
		user.DestroyAllOtherSessions(c)
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at": "controllers.UpdateUserName",
		}).Info("user password updated")
	}
}

func deleteUser(c *gin.Context) {
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}
	admin, err := currentUser(c)
	if err != nil {
		return
	}

	username, ok := user_info["username"]
	if !ok {
		err := "no old password provided"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.deleteUser",
			"error": err,
			"admin": admin.Username,
		}).Error("error updating user password")
		return
	}

	user, err := models.UserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.deleteUser",
			"error": err.Error(),
			"admin": admin.Username,
		}).Error("error looking up user to delete")
		return
	}

	only_admin, err := user.OnlyAdmin()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.deleteUser",
			"error": err.Error(),
			"admin": admin.Username,
		}).Error("error checking admin status of user to delete")
		return
	}
	if only_admin {
		err := "user is only remaining admin"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.deleteUser",
			"error": err,
			"admin": admin.Username,
		}).Error("error updating user password")
		return
	}

	user.DestroyAllSessions()
	err = user.Delete()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.deleteUser",
			"error": err.Error(),
			"admin": admin.Username,
		}).Error("error deleting user")
		return
	} else {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at":       "controllers.deleteUser",
			"username": username,
			"admin":    admin.Username,
		}).Info("deleted user")
	}

}
