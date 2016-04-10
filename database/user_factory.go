package database

func TestUser(traits []string) (user User) {
	user = User{
		Name:  "Turd Ferguson",
		Email: "bighat@example.com",
	}
	// apply each trait
	DB().Create(&user)
	return user
}
