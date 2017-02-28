package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Session struct {
	gorm.Model
	UserID            uint
	Cookie            string
	OriginallyCreated time.Time
	LastUsed          time.Time
}

func (session Session) HTTPCookie() http.Cookie {
	cookie := http.Cookie{
		Name:     "wave_session",
		Value:    session.Cookie,
		Path:     "/",
		Domain:   helpers.WaveHostname,
		MaxAge:   int(time.Now().AddDate(1, 0, 1).Unix()),
		Secure:   helpers.TLS,
		HttpOnly: true,
		Raw:      "wave_session=" + session.Cookie,
		Unparsed: []string{"wave_session=" + session.Cookie},
	}

	return cookie
}

func (session Session) User() (user User, err error) {
	db_err := Orm.Model(&session).Related(&user)
	err = db_err.Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "(Session) models.User",
			"error": err.Error(),
		}).Warn("error finding related user for session")
	}
	return
}

func (session *Session) Save() error {
	return Orm.Save(&session).Error
}

func (session *Session) Reload() error {
	return Orm.First(&session, "Cookie = ?", session.Cookie).Error
}

func (session *Session) Delete() error {
	return Orm.Delete(&session).Error
}

func SessionFromID(id string) (session Session, err error) {
	db_err := Orm.First(&session, "Cookie = ?", id)
	err = db_err.Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.SessionFromID",
			"error": err.Error(),
		}).Warn("error looking up session")
	} else {
		session.LastUsed = time.Now()
		err = session.Save()
	}
	return
}
