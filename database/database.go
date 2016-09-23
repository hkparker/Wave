package database

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
)

var Orm *gorm.DB

func Connect(db_username, db_password, db_name, db_ssl string) {
	var db_type string
	var db_args string
	if helpers.Production() || helpers.Development() {
		db_type = "postgres"
		db_args = fmt.Sprintf(
			"user=%s password=%s sslmode=%s",
			db_username,
			db_password,
			db_name,
			db_ssl,
		)
	} else if helpers.Testing() {
		db_type = "sqlite3"
		db_args = ":memory:"
	} else {
		log.WithFields(log.Fields{
			"environment": helpers.Env(),
		}).Fatal("unknown database environment")
	}
	var err error
	Orm, err = gorm.Open(db_type, db_args)
	if err != nil {
		log.WithFields(log.Fields{
			"type":  db_type,
			"user":  db_username,
			"ssl":   db_ssl,
			"error": err.Error(),
		}).Fatal("unable to connect to database server")
	}

	// create the table if missing?

	if helpers.Testing() {
		Init()
	}
}

func Init() {
	Orm.CreateTable(User{})
	Orm.CreateTable(Session{})
}
