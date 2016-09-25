package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Name               string
	Password           []byte
	Email              string `sql:"not null;unique"`
	Admin              bool
	Sessions           []Session
	PasswordResetToken string
}

func init() {
	if !database.Orm.HasTable(User{}) {
		database.Orm.CreateTable(User{})
	}
}

func CreateUser(email string) (err error) {
	user := User{
		Email:              email,
		PasswordResetToken: helpers.RandomString(),
	}
	db_err := database.Orm.Create(&user)
	if db_err.Error != nil {
		err = db_err.Error
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("error_saving_user")
	} else {
		// user.EmailAccountSetup()
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"email":  "account_register",
		}).Info("email_sent")
	}
	log.WithFields(log.Fields{
		"UserID": user.ID,
		"email":  user.Email,
	}).Info("user_created")
	return
}

func CurrentUser(c *gin.Context) (user User, err error) {
	session_cookie, err := c.Request.Cookie("wave_session")
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "database.CurrentUser",
			"error": err.Error(),
		}).Error("session_missing")
		return
	}

	session, err := SessionFromID(session_cookie.Value)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "database.CurrentUser",
			"error": err,
		}).Error("session_missing")
		return
	}
	return session.ActiveUser()
}

func (user *User) SetPassword(password string) (err error) {
	pw_data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("bcrypt_error")
		return
	}
	user.Password = pw_data
	database.Orm.Save(&user)
	// err
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_set")
	return
}

func (user *User) ResetPassword() (err error) {
	user.DestroyAllSessions()
	user.PasswordResetToken = helpers.RandomString()
	db_err := database.Orm.Save(&user)
	if db_err.Error != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
		}).Warn("error_saving_user")
		err = db_err.Error
	} else {
		// user.EmailPasswordReset()
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"email":  "password_reset",
		}).Info("email_sent")
	}
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_reset")
	return
}

func (user *User) NewSession() (wave_session string, err error) {
	now := time.Now()
	wave_session = helpers.RandomString()
	session := Session{
		UserID:            user.ID,
		OriginallyCreated: now,
		LastUsed:          now,
		Cookie:            wave_session,
	}
	user.Sessions = append(user.Sessions, session)
	database.Orm.Save(&user)
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("session_created")
	return
}

func (user *User) DestroyAllSessions() {
	user.Sessions = []Session{}
	database.Orm.Save(&user)
	for session := range user.Sessions {
		database.Orm.Unscoped().Delete(&session)
	}
}

func (user *User) Reload() {
	database.Orm.First(&user, "Email = ?", user.Email)
}
