package models

import (
	log "github.com/sirupsen/logrus"
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
	db_err := Orm.Create(&user).Error
	if db_err != nil {
		err = db_err
		log.WithFields(log.Fields{
			"at":     "models.CreateUser",
			"UserID": user.ID,
			"error":  err,
		}).Warn("error_saving_user")
	} else {
		db_err = user.Save()
		if db_err != nil {
			log.WithFields(log.Fields{
				"at":     "models.CreateUser",
				"UserID": user.ID,
			}).Warn("error_saving_user")
			err = db_err
			return
		} else {
			log.WithFields(log.Fields{
				"at":       "models.CreateUser",
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
			"at":     "models.SetPassword",
			"UserID": user.ID,
			"error":  err,
		}).Warn("bcrypt_error")
		return
	}
	user.Password = pw_data
	err = user.Save()
	if err != nil {
		log.WithFields(log.Fields{
			"at":     "models.SetPassword",
			"UserID": user.ID,
			"error":  err.Error(),
		}).Error("error_setting_password")
	} else {
		log.WithFields(log.Fields{
			"at":     "models.SetPassword",
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
			"at":     "models.NewSession",
			"UserID": user.ID,
			"error":  err.Error(),
		}).Error("error_creating_sessions")
	} else {
		log.WithFields(log.Fields{
			"at":     "models.NewSession",
			"UserID": user.ID,
		}).Info("session_created")
	}
	return
}

func (user *User) SessionCount() int {
	return Orm.Model(&user).Association("Sessions").Count()
}

func (user *User) DestroyAllOtherSessions(session_cookie string) {
	var new_sessions []Session
	var sessions []Session
	Orm.Model(&user).Related(&sessions)
	for _, session := range sessions {
		if session.Cookie != session_cookie {
			Orm.Model(&user).Association("Sessions").Delete(&session)
			Orm.Unscoped().Delete(&session)
		} else {
			new_sessions = append(new_sessions, session)
		}
	}
	user.Sessions = new_sessions
	user.Save()
}

func (user *User) DestroyAllSessions() {
	var sessions []Session
	Orm.Model(&user).Related(&sessions)
	for _, session := range sessions {
		Orm.Model(&user).Association("Sessions").Delete(&session)
		Orm.Unscoped().Delete(&session)
	}
	user.Sessions = []Session{}
	user.Save()
}

func (user User) OnlyAdmin() (only_admin bool, err error) {
	only_admin = false
	if !user.Admin {
		return
	}

	var admins []User
	err = Orm.Where("Admin = ?", true).Find(&admins).Error
	if err != nil {
		return
	}

	if len(admins) == 1 {
		only_admin = true
	}
	return
}

func (user *User) Reload() error {
	return Orm.Unscoped().First(&user, "Username = ?", user.Username).Error
}

func (user *User) Save() error {
	return Orm.Save(&user).Error
}

func (user *User) Delete() error {
	return Orm.Delete(&user).Error
}

func UserByUsername(username string) (user User, err error) {
	err = Orm.First(&user, "Username = ?", username).Error
	return
}

func UserFromSessionCookie(session_cookie string) (user User, err error) {
	session, err := SessionFromID(session_cookie)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.UserFromSessionCookie",
			"error": err,
		}).Error("session_missing")
		return
	}
	user, err = session.User()
	return
}

func Users() (users []User, err error) {
	err = Orm.Find(&users).Error
	return
}

// Used for tests
func DropUsers() {
	Orm.Delete(User{})
}
