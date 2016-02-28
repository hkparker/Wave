package helpers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"os"
)

var pg *gorm.DB

func SetupPostgres() {
	db, err := gorm.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to postgres")
		os.Exit(1)
	}
	pg = &db
}

func DB() *gorm.DB {
	if pg == nil {
		SetupPostgres()
	}
	return pg
}
