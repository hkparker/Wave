package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/database"
)

func init() {
	if database.Orm != nil {
		CreateTables()
	}
}

func CreateTables() {
	if !database.Orm.HasTable(User{}) {
		database.Orm.CreateTable(User{})
		log.Info("creating missing user table")
	}

	if !database.Orm.HasTable(Session{}) {
		database.Orm.CreateTable(Session{})
		log.Info("creating missing session table")
	}

	if !database.Orm.HasTable(Collector{}) {
		database.Orm.CreateTable(Collector{})
		log.Info("creating missing collector table")
	}
}
