package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Orm *gorm.DB

func Connect() {
	if helpers.DBAdminUsername != "" {
		db_check_args := fmt.Sprintf(
			"user=%s password=%s sslmode=%s",
			helpers.DBAdminUsername,
			helpers.DBAdminPassword,
			helpers.DBTLS,
		)
		var err error
		check, err := gorm.Open("postgres", db_check_args)
		if err != nil {
			log.WithFields(log.Fields{
				"at":    "models.Connect",
				"user":  helpers.DBAdminUsername,
				"ssl":   helpers.DBTLS,
				"error": err.Error(),
			}).Fatal("unable to connect to database server")
		}
		if check.Exec(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", helpers.DBName)).RowsAffected != 1 {
			check.Exec(fmt.Sprintf("CREATE DATABASE %s", helpers.DBName))
			log.WithFields(log.Fields{
				"at":      "models.Connect",
				"db_name": helpers.DBName,
			}).Info("created missing database")
		}
	}

	db_args := fmt.Sprintf(
		"user=%s password=%s sslmode=%s dbname=%s",
		helpers.DBUsername,
		helpers.DBPassword,
		helpers.DBTLS,
		helpers.DBName,
	)
	var err error
	Orm, err = gorm.Open("postgres", db_args)
	if err != nil {
		log.WithFields(log.Fields{
			"at":     "models.Connect",
			"user":   helpers.DBUsername,
			"ssl":    helpers.DBTLS,
			"dbname": helpers.DBName,
			"error":  err.Error(),
		}).Fatal("unable to connect to database server")
	}
	Setup()
}
