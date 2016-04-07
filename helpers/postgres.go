package helpers

import (
	log "github.com/Sirupsen/logrus"
	_ "github.com/erikstmartin/go-testdb"
	"github.com/jinzhu/gorm"
	"os"
)

var pg *gorm.DB
var testpg *gorm.DB

// InitTestDB  // connect to the test database and create all tables
func TestDB() *gorm.DB {
	if testpg == nil {
		db, err := gorm.Open("testdb", "")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("unable to connect to postgres")
			os.Exit(1)
		}
		testpg = &db
	}
	return testpg
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
		pg = &db
	}
	return pg
}
