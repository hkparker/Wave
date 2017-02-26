package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
)

func init() {
	if database.Orm != nil {
		createTables()
	}
}

func Setup() {
	createTables()
	createAdmin()
}

func createTables() {
	if !database.Orm.HasTable(User{}) {
		database.Orm.CreateTable(User{})
		log.Info("creating missing user table")
	}

	if !database.Orm.HasTable(Session{}) {
		database.Orm.CreateTable(Session{})
		log.Info("creating missing session table")
	}

	if !database.Orm.HasTable(Collector{}) {
		database.Orm.CreateTable(Collector{})
		log.Info("creating missing collector table")
	}

	if !database.Orm.HasTable(TLS{}) {
		database.Orm.CreateTable(TLS{})
		log.Info("creating missing tls configuration table")
	}
}

func createAdmin() {
	var admins []User
	if err := database.Orm.Where("Admin = ?", true).Find(&admins).Error; err == nil {
		if len(admins) == 0 {
			var user User
			err = database.Orm.First(&user, "Username = ?", "root").Error
			if err == nil {
				database.Orm.Unscoped().Delete(&user)
			}
			admin := User{
				Username: "root",
				Admin:    true,
			}
			password := helpers.RandomString()
			err = admin.SetPassword(password)
			if err != nil {
				log.Fatal(err)
			}
			err = admin.Save()
			if err != nil {
				log.Fatal(err)
			}
			log.WithFields(log.Fields{
				"username": "root",
				"password": password,
			}).Info("created_default_admin")
		}
	} else {
		log.Fatal(err)
	}
}
