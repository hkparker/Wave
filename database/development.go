package database

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ReseedDevelopmentDatabase() {

}

func createDevelopmentDatabase() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "user=postgres dbname=wave_dev sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to development postgres database")
	}
	return
}

func developmentDB() *gorm.DB {
	if developmentdb == nil {
		developmentdb = createDevelopmentDatabase()
	}
	return developmentdb
}
