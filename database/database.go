package database

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
)

var productiondb *gorm.DB
var developmentdb *gorm.DB
var testdb *gorm.DB

func DB() *gorm.DB {
	if helpers.Production() {
		return productionDB()
	} else if helpers.Development() {
		return developmentDB()
	} else if helpers.Testing() {
		return testDB()
	}
	log.WithFields(log.Fields{
		"environment": helpers.Env(),
	}).Fatal("unknown database environment")
	return nil
}
