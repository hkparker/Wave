package models

import (
	"github.com/hkparker/Wave/helpers"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func skipTestNewSessionCreatesSession(t *testing.T) {
	user := User{}
	helpers.TestDB().Where(User{Email: "usertest@example.com"}).First(&user)
	cookie, err := user.NewSession()
	assert.Nil(t, err)
	session := Session{}
	helpers.TestDB().Where(Session{Cookie: cookie}).First(&session)
	assert.Equal(t, session.UserID, user.ID)
}

func TestSessionCreatedWhenLoggedIn(t *testing.T) {

}
