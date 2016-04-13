package database

import (
	"crypto/rand"
	log "github.com/Sirupsen/logrus"
	"github.com/jbenet/go-base58"
)

func TestUser(traits []string) (user User) {
	email_bytes := make([]byte, 32)
	rand.Read(email_bytes)
	email := base58.Encode(email_bytes) + "@example.com"
	log.WithFields(log.Fields{
		"traits": traits,
	}).Info("creating test user")
	user = User{
		Name:  "Turd Ferguson",
		Email: email,
	}

	for _, trait := range traits {
		switch trait {
		case "admin":
			user.Name = "Wifi Jackson"
			user.Admin = true
		}
	}

	DB().Create(&user)
	return
}
