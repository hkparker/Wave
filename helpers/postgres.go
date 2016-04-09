package helpers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"os"
)

var pg *gorm.DB
var testdb *gorm.DB

func createTestDatabase() (db *gorm.DB) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to sqlite")
	}
	return
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

func ReseedDevelopmentDatabse() {

}

func createProductionDatabase() {

}

func TestDB() *gorm.DB {
	if testdb == nil {
		testdb = createTestDatabase()
	}
	return testdb
}

func DB() *gorm.DB {
	if pg == nil {
		db, err := gorm.Open("postgres", "user=postgres dbname=wave sslmode=disable")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("unable to connect to postgres")
			os.Exit(1)
		}
		pg = db
	}
	return pg
}
