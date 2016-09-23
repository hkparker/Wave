package database

import (
	"github.com/hkparker/Wave/helpers"
)

func TestUser(traits []string) (user User) {
	email := helpers.RandomString() + "@example.com"
	user = User{
		Name:  "Turd Ferguson",
		Email: email,
	}
	user.SetPassword(helpers.RandomString())

	for _, trait := range traits {
		switch trait {
		case "admin":
			user.Name = "Wifi Jackson"
			user.Admin = true
		}
	}

	//	DB().Create(&user)
	return
}
