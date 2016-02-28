package models

import (
	"flag"
	"github.com/hkparker/Wave/helpers"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func seedUserTests() {
	helpers.SetupPostgres()
	helpers.DB().CreateTable(&User{})
	helpers.DB().CreateTable(&Session{})
	user := User{
		Name:     "Joe Hackerman",
		Password: []byte{},
		Email:    "usertest@example.com",
	}
	helpers.DB().Create(&user)

}

func TestNewSessionCreatesSession(t *testing.T) {
	user := User{}
	helpers.DB().Where(User{Email: "usertest@example.com"}).First(&user)
	cookie, err := user.NewSession()
	if err != nil {
		t.Fatal("error creating session:", err)
	}
	session := Session{}
	helpers.DB().Where(Session{Cookie: cookie}).First(&session)
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
