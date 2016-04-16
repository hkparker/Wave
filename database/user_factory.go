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

	for _, trait := range traits {
		switch trait {
		case "admin":
			user.Name = "Wifi Jackson"
			user.Admin = true
		case "with_password_reset":
			user.ResetPassword()
		}
	}

	DB().Create(&user)
	return
}
