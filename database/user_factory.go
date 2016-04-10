package database

func TestUser(traits []string) (user User) {
	user = User{
		Name:  "Turd Ferguson",
		Email: "bighat@example.com",
	}

	for _, trait := range traits {
		switch trait {
		case "admin":
			user.Admin = true
		}
	}

	DB().Create(&user)
	return user
}
