package models

import (
	"flag"
	"github.com/hkparker/Wave/helpers"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, err)
	session := Session{}
	helpers.DB().Where(Session{Cookie: cookie}).First(&session)
	assert.Equal(t, session.UserID, user.ID)
}

func TestSessionCreatedWhenLoggedIn(t *testing.T) {

}

func TestMain(m *testing.M) {
	flag.Parse()
	// database setup and teardown
	os.Exit(m.Run())
}
