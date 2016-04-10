package models

import (
	"github.com/hkparker/Wave/database"
)

func TestUser(traits []string) (user User) {
	user = User{
		Name:  "Turd Ferguson",
		Email: "bighat@example.com",
	}
	// apply each trait
	database.TestDB().Create(&user)
	return user
}
