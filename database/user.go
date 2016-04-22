package database

import (
	log "github.com/Sirupsen/logrus"
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

func CreateUser(email string) (err error) {
	user := User{
		Email: email,
	}
	db_err := DB().Create(&user)
	if db_err.Error != nil {
		err = db_err.Error
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("error_saving_user")
	} else {
		// user.AccountSetupEmail() // reset password and email the link with a welcome message
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
	DB().Save(&user)
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_set")
	return
}

func (user *User) ResetPassword() (err error) {
	user.DestroyAllSessions()
	user.PasswordResetToken = helpers.RandomString()
	db_err := DB().Save(&user)
	if db_err.Error != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
		}).Warn("error_saving_user")
		err = db_err.Error
	} else {
		//user.EmailPasswordReset(user.PasswordResetToken)
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

func ValidAuthentication(email, password string) (valid bool) {
	valid = false

	var user User
	db_err := DB().First(&user, "Email = ?", email)
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
	DB().Save(&user)
	log.WithFields(log.Fields{
		"UserID": user.ID,
		"Cookie": wave_session,
	}).Info("session_created")
	return
}

func (user *User) DestroyAllSessions() {
	user.Sessions = []Session{}
	DB().Save(&user)
	for session := range user.Sessions {
		DB().Unscoped().Delete(&session)
	}
}
