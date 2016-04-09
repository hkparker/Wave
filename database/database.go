package database

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

var productiondb *gorm.DB
var developmentdb *gorm.DB
var testdb *gorm.DB

func ReseedDevelopmentDatabase() {

}

func TestDB() *gorm.DB {
	if testdb == nil {
		testdb = createTestDatabase()
	}
	return testdb
}
func DB() *gorm.DB {
	if helpers.Production() {
		return productionDB()
	} else if helpers.Development() {
		return developmentDB()
	}
	log.WithFields(log.Fields{
		"environment": helpers.Env(),
	}).Fatal("unknown database environment")
	return nil
}
