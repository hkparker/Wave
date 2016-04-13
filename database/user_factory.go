package database

import (
	log "github.com/Sirupsen/logrus"
)

func TestUser(traits []string) (user User) {
	log.WithFields(log.Fields{
		"traits": traits,
	}).Info("creating test user")
	user = User{
		Name:  "Turd Ferguson",
		Email: "bighat@example.com",
	}

	for _, trait := range traits {
		switch trait {
		case "admin":
			user.Name = "Wifi Jackson"
			user.Email = "ughbluergna@murf.mnrgg"
			user.Admin = true
		}
	}

	DB().Create(&user)
	return
}
