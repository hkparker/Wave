package helpers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"os"
)

var env = "production"
var pg *gorm.DB
var testpg *gorm.DB

func SetEnv(newenv string) {
	env = newenv
}

func DB() *gorm.DB {
	switch env {
	case "testing":
		return TestDB()
	}
	return ProductionDB()
}

func TestDB() *gorm.DB {
	if testpg == nil {
		db, err := gorm.Open("postgres", "user=postgres dbname=wave_test sslmode=disable")
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

func ProductionDB() *gorm.DB {
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
