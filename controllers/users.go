package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
)

// Handle a request from administrator to create a new user
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

	// Save the new user and create a link for the user
	// to set a password
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
		// Report if the user could not be created
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.createUser",
			"username": username,
			"error":    err.Error(),
			"admin":    admin.Username,
		}).Error("error creating user")
	}
}

// Handle requests for users to change their own names
func updateUserName(c *gin.Context) {
	// Ensure the request is valid JSON and get
	// the user from the current session
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}
	user, err := currentUser(c)
	if err != nil {
		return
	}

	// Ensure the JSON contains a username key with data
	username, ok := user_info["username"]
	if !ok || username == "" {
		name_error := "no username provided"
		c.JSON(500, gin.H{"error": name_error})
		log.WithFields(log.Fields{
			"at":       "controllers.updateUserName",
			"error":    name_error,
			"username": username,
			"user":     user.Name,
		}).Error("error updating user name")
		return
	}

	// Update the users name and save to the database
	user.Name = username
	db_err := user.Save()
	if db_err != nil {
		// Return an error if the update could not be saved
		c.JSON(500, gin.H{"error": db_err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.updateUserName",
			"error":    db_err.Error(),
			"username": username,
			"user":     user.Name,
		}).Error("error updating user name")
	} else {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at":       "controllers.updateUserName",
			"username": username,
			"old_name": user.Name,
		}).Info("user name updated")
	}
}

// Handle requests for password changes from users
func updateUserPassword(c *gin.Context) {
	// Ensure the request is valid JSON and get
	// the user from the current session
	user_info, err := requestJSON(c)
	if err != nil {
		return
	}
	user, err := currentUser(c)
	if err != nil {
		return
	}

	// Ensure the JSON contains an old_password key
	old_password, ok := user_info["old_password"]
	if !ok {
		err := "no old password provided"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.updateUserPassword",
			"error": err,
			"user":  user.Name,
		}).Error("error updating user password")
		return
	}

	// Ensure the JSON contains an new_password key
	new_password, ok := user_info["new_password"]
	if !ok {
		err := "no new password provided"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.updateUserPassword",
			"error": err,
			"user":  user.Name,
		}).Error("error updating user password")
		return
	}

	// Ensure the old password is valid authentication
	if !user.ValidAuthentication(old_password) {
		err := "old password incorrect"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":    "controllers.updateUserPassword",
			"error": err,
			"user":  user.Name,
		}).Error("error updating user password")
		return
	}

	// Set the new password for the user
	err = user.SetPassword(new_password)
	if err != nil {
		// Report an error if the user could not be saved
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.updateUserPassword",
			"error": err.Error(),
			"user":  user.Name,
		}).Error("error setting user password")
	} else {
		// Invalidate all session except the one which made the request
		user.DestroyAllOtherSessions(c)
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at":   "controllers.updateUserPassword",
			"user": user.Name,
		}).Info("user password updated")
	}
}

// Handle request from administrator to delete a user
func deleteUser(c *gin.Context) {
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
		err := "no username provided"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":       "controllers.deleteUser",
			"username": username,
			"error":    err,
			"admin":    admin.Username,
		}).Error("error updating user password")
		return
	}

	// Ensure the user exists
	user, err := models.UserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.deleteUser",
			"username": username,
			"error":    err.Error(),
			"admin":    admin.Username,
		}).Error("error looking up user to delete")
		return
	}

	// Ensure the user is not the only remaining administrator removing themselves
	only_admin, err := user.OnlyAdmin()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.deleteUser",
			"username": username,
			"error":    err.Error(),
			"admin":    admin.Username,
		}).Error("error checking admin status of user to delete")
		return
	}
	if only_admin {
		err := "user is only remaining admin"
		c.JSON(500, gin.H{"error": err})
		log.WithFields(log.Fields{
			"at":       "controllers.deleteUser",
			"username": username,
			"error":    err,
			"admin":    admin.Username,
		}).Error("error updating user password")
		return
	}

	// Destroy all sessions and delete the user
	user.DestroyAllSessions()
	err = user.Delete()
	if err != nil {
		// Return an error if the user could not be deleted
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":       "controllers.deleteUser",
			"username": username,
			"error":    err.Error(),
			"admin":    admin.Username,
		}).Error("error deleting user")
	} else {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at":       "controllers.deleteUser",
			"username": username,
			"admin":    admin.Username,
		}).Info("deleted user")
	}

}
