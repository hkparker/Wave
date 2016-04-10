package database

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func createProductionDatabase() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "user=postgres dbname=wave sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to production postgres database")
	}
	return
}

func productionDB() *gorm.DB {
	if productiondb == nil {
		productiondb = createProductionDatabase()
	}
	return productiondb
}
