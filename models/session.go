package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/database"
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID            uint
	Cookie            string
	OriginallyCreated time.Time
	LastUsed          time.Time
}

func init() {
	if !database.Orm.HasTable(Session{}) {
		database.Orm.CreateTable(Session{})
	}
}

func SessionFromID(id string) (session Session, err error) {
	db_err := database.Orm.First(&session, "Cookie = ?", id)
	err = db_err.Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":         "database.SessionFromID",
			"session_id": id,
			"error":      err.Error(),
		}).Warn("error looking up session")
	}
	return
}

func (session Session) ActiveUser() (user User, err error) {
	db_err := database.Orm.Model(&session).Related(&user)
	err = db_err.Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "database.Session.Active",
			"error": err.Error(),
		}).Warn("error finding related user for session")
	}
	return
}
