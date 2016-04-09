package models

func CreateUser(traits []string) (user User) {
	return User{
		Name:  "Turd Ferguson",
		Email: "bighat@example.com",
	}
}
