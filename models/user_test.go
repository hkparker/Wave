package models

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func seedUserTests() gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	SetDB(db)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to postgres")
	}
	db.CreateTable(&User{})
	db.CreateTable(&Session{})
	user := User{
		Name:     "Joe Hackerman",
		Password: []byte{},
		Email:    "usertest@example.com",
	}
	db.Create(&user)
	return db

}

func TestNewSessionCreatesSession(t *testing.T) {
	db := seedUserTests()
	user := User{}
	db.Where(User{Email: "usertest@example.com"}).First(&user)
	cookie, err := user.NewSession()
	if err != nil {
		t.Fatal("error creating session:", err)
	}
	session := Session{}
	db.Where(Session{Cookie: cookie}).First(&session)
	if session.UserID != user.ID {
		t.Fatal("NewSession did not create a session with proper information")
	}
}

func TestSessionCreatedWhenLoggedIn(t *testing.T) {

}

func TestMain(m *testing.M) {
	flag.Parse()
	// database setup and teardown
	os.Exit(m.Run())
}
