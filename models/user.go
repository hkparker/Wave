package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Password []byte
	Username string `sql:"not null;unique"`
	Admin    bool
	Sessions []Session
}

func CreateUser(username string) (err error) {
	user := User{
		Username: username,
	}
	db_err := database.Orm.Create(&user).Error
	if db_err != nil {
		err = db_err
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("error_saving_user")
	} else {
		db_err = user.Save()
		if db_err != nil {
			log.WithFields(log.Fields{
				"UserID": user.ID,
			}).Warn("error_saving_user")
			err = db_err
			return
		} else {
			log.WithFields(log.Fields{
				"UserID":   user.ID,
				"username": user.Username,
			}).Info("user_created")
		}
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
	err = user.Save()
	if err != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err.Error(),
		}).Error("error_setting_password")
	} else {
		log.WithFields(log.Fields{
			"UserID": user.ID,
		}).Info("user_password_set")
	}
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
	err = user.Save()
	if err != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err.Error(),
		}).Error("error_creating_sessions")
	} else {
		log.WithFields(log.Fields{
			"UserID": user.ID,
		}).Info("session_created")
	}
	return
}

func (user *User) DestroyAllOtherSessions(session_cookie string) {
	var sessions []Session
	for _, session := range user.Sessions {
		if session.Cookie != session_cookie {
			err := database.Orm.Unscoped().Delete(&session)
			if err != nil {

			}
		} else {
			sessions = append(sessions, session)
		}
	}
	user.Sessions = sessions
	user.Save()
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
	only_admin = false
	if !user.Admin {
		return
	}

	var admins []User
	err = database.Orm.Where("Admin = ?", true).Find(&admins).Error
	if err != nil {
		return
	}

	if len(admins) == 1 {
		only_admin = true
	}
	return
}

func (user *User) Reload() error {
	return database.Orm.Unscoped().First(&user, "Username = ?", user.Username).Error
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

// Used for tests
func DropUsers() {
	database.Orm.Delete(User{})
}
