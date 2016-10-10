package database

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Orm *gorm.DB

func init() {
	if helpers.TestingCmd() && Orm == nil {
		var err error
		Orm, err = gorm.Open("sqlite3", ":memory:")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("unable to connect to testing database server")
		}
	}
}

func Connect(db_username, db_password, db_name, db_ssl string) {
	db_args := fmt.Sprintf(
		"user=%s password=%s sslmode=%s",
		db_username,
		db_password,
		db_ssl,
		//db_name,
	)
	var err error
	Orm, err = gorm.Open("postgres", db_args)
	if err != nil {
		log.WithFields(log.Fields{
			"user":  db_username,
			"ssl":   db_ssl,
			"error": err.Error(),
		}).Fatal("unable to connect to database server")
	}
}
