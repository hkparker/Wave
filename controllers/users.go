package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

	err = models.CreateUser(email)
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

func UpdateUserName(c *gin.Context) {
	var user_info map[string]string
	err := c.BindJSON(&user_info)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserName",
			"error": err,
		}).Error("error parsing request")
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

	user, err := models.CurrentUser(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserName",
			"error": err.Error,
		}).Error("error getting current user")
		return
	}

	user.Name = name
	db_err := database.Orm.Save(&user)
	if db_err.Error != nil {
		c.JSON(500, gin.H{"error": db_err.Error})
		log.WithFields(log.Fields{
			"at":    "controllers.UpdateUserName",
			"error": db_err.Error,
		}).Error("error updating user")
	} else {
		c.JSON(200, gin.H{"success": "true"})
		log.WithFields(log.Fields{
			"at": "controllers.UpdateUserName",
		}).Info("user name updated")
	}
}

func Login(c *gin.Context) {
	user, err := models.CurrentUser(c)
	if err == nil {
		wave_session, err := user.NewSession()
		if err == nil {
			http.SetCookie(
				c.Writer,
				&http.Cookie{
					Name:  "wave_session",
					Value: wave_session,
				},
			)
		} else {
			// log
		}
	} else {
		// unauthorized
	}
}

func PasswordReset(c *gin.Context) {

}

func validAuthentication(email, password string) (valid bool) {
	valid = false

	var user models.User
	db_err := database.Orm.First(&user, "Email = ?", email)
	if db_err.Error != nil {
		return
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return
	}

	valid = true
	return
}

//func validateNewUser(c *gin.Context) (map[string]string, bool) {
//	var user_info map[string]string
//	err := c.BindJSON(&user_info)
//	if err != nil {
//		c.JSON(500, gin.H{"error": err.Error()})
//		log.WithFields(log.Fields{
//			"at":    "controllers.CreateUser",
//			"error": err,
//		}).Error("error parsing request")
//		return user_info, false
//	}
//
//	_, ok := user_info["email"]
//	if !ok {
//		email_error := "no email provided"
//		c.JSON(500, gin.H{"error": email_error})
//		log.WithFields(log.Fields{
//			"at":    "controllers.CreateUser",
//			"error": email_error,
//		}).Error("error creating user")
//		return user_info, false
//	}
//
//	// ensure uniqueness of email
//	return user_info, true
//}
