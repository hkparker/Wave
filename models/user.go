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
	Username           string `sql:"not null;unique"`
	Admin              bool
	Sessions           []Session
	PasswordResetToken string
	//PasswordResetTime
}

func CreateUser(username string) (password_reset_link string, err error) {
	user := User{
		Username:           username,
		PasswordResetToken: helpers.RandomString(),
	}
	db_err := database.Orm.Create(&user).Error
	if db_err != nil {
		err = db_err
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("error_saving_user")
	} else {
		log.WithFields(log.Fields{
			"UserID":   user.ID,
			"username": user.Username,
		}).Info("user_created")

		password_reset_token := helpers.RandomString()
		user.PasswordResetToken = password_reset_token
		db_err = user.Save()
		if db_err != nil {
			log.WithFields(log.Fields{
				"UserID": user.ID,
			}).Warn("error_saving_user")
			err = db_err
			return
		}
		password_reset_link = helpers.WaveAddress + "/users/reset/" + password_reset_token
	}
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
	user.Save()
	// err
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_set")
	return
}

func (user *User) ResetPassword() (err error) {
	user.DestroyAllSessions()
	user.SetPassword(helpers.RandomString())
	user.PasswordResetToken = helpers.RandomString()
	err = user.Save()
	if err != nil {
		log.WithFields(log.Fields{
			"at":     "user.ResetPassword",
			"UserID": user.ID,
		}).Warn("error_saving_user")
		return
	}
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_reset")
	return
}

func (user User) ValidAuthentication(password string) (valid bool) {
	valid = false

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
	user.Save()
	// err
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("session_created")
	return
}

func (user *User) DestroyAllOtherSessions(c *gin.Context) {

}

func (user *User) DestroyAllSessions() {
	user.Sessions = []Session{}
	err := user.Save()
	if err != nil {

	}
	for session := range user.Sessions {
		err := database.Orm.Unscoped().Delete(&session)
		if err != nil {

		}
	}
}

func (user User) OnlyAdmin() (only_admin bool, err error) {
	only_admin = true

	var admins []User
	err = database.Orm.Where("Admin = ?", true).Find(&admins).Error
	if err != nil {
		return
	}

	if len(admins) > 1 || !user.Admin {
		only_admin = false
	}
	return
}

func (user *User) Reload() error {
	return database.Orm.First(&user, "Username = ?", user.Username).Error
}

func (user *User) Save() error {
	return database.Orm.Save(&user).Error
}

func (user *User) Delete() error {
	return database.Orm.Delete(&user).Error
}

func UserByUsername(username string) (user User, err error) {
	err = database.Orm.First(&user, "Username = ?", username).Error
	return
}

func UserFromSessionCookie(session_cookie string) (user User, err error) {
	session, err := SessionFromID(session_cookie)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "database.currentUser",
			"error": err,
		}).Error("session_missing")
		return
	}
	user, err = session.User()
	return
}

func Users() (users []User, err error) {
	err = database.Orm.Find(&users).Error
	return
}
