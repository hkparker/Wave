package models

import (
	"github.com/hkparker/Wave/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSessionCreatesSession(t *testing.T) {
	user := CreateUser([]string{})
	cookie, err := user.NewSession()
	assert.Nil(t, err)
	session := Session{}
	helpers.TestDB().Where(Session{Cookie: cookie}).First(&session)
	if assert.NotNil(t, session.UserID) {
		assert.Equal(t, session.UserID, user.ID)
	}
}

func TestSessionCreatedWhenLoggedIn(t *testing.T) {

}
