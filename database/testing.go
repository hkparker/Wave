package database

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

func createTestDatabase() (db *gorm.DB) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("error creating in-memory sqlite testing database")
	}
	return
}
