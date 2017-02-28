package models

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Orm *gorm.DB

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
			"at":    "models.Connect",
			"user":  db_username,
			"ssl":   db_ssl,
			"error": err.Error(),
		}).Fatal("unable to connect to database server")
	}
	Setup()
}
