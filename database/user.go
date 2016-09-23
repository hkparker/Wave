package database

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
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

//func CreateUser(email string) (err error) {
//	user := User{
//		Email:              email,
//		PasswordResetToken: helpers.RandomString(),
//	}
//	db_err := DB().Create(&user)
//	if db_err.Error != nil {
//		err = db_err.Error
//		log.WithFields(log.Fields{
//			"UserID": user.ID,
//			"error":  err,
//		}).Warn("error_saving_user")
//	} else {
//		// user.EmailAccountSetup()
//		log.WithFields(log.Fields{
//			"UserID": user.ID,
//			"email":  "account_register",
//		}).Info("email_sent")
//	}
//	log.WithFields(log.Fields{
//		"UserID": user.ID,
//		"email":  user.Email,
//	}).Info("user_created")
//	return
//}

func CurrentUser(c *gin.Context) (user User, err error) {
	sess_ids, present := c.Request.Header["wave_session"]
	if !present {
		err = errors.New("no wave session header")
		log.WithFields(log.Fields{
			"at":    "database.CurrentUser",
			"error": err,
		}).Error("session_missing")
		return
	}
	if len(sess_ids) != 1 {
		err = errors.New("")
		return
	}
	session, err := SessionFromID(sess_ids[0])
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
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_set")
	return
}

func (user *User) ResetPassword() (err error) {
	//user.DestroyAllSessions()
	user.PasswordResetToken = helpers.RandomString()
	//if db_err.Error != nil {
	//	log.WithFields(log.Fields{
	//		"UserID": user.ID,
	//	}).Warn("error_saving_user")
	//	err = db_err.Error
	//} else {
	//	// user.EmailPasswordReset()
	//	log.WithFields(log.Fields{
	//		"UserID": user.ID,
	//		"email":  "password_reset",
	//	}).Info("email_sent")
	//}
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
	log.WithFields(log.Fields{
		"UserID": user.ID,
		"Cookie": wave_session,
	}).Info("session_created")
	return
}

//func (user *User) DestroyAllSessions() {
//	user.Sessions = []Session{}
//	DB().Save(&user)
//	for session := range user.Sessions {
//		DB().Unscoped().Delete(&session)
//	}
//}
