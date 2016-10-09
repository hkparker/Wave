package models

import (
	"github.com/hkparker/Wave/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionFromIDLoadsValidSession(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	cookie, err := user.NewSession()
	assert.Nil(err)

	session, err := SessionFromID(cookie)
	assert.Nil(err)

	var related User
	db_err := database.Orm.Model(&session).Related(&related)
	assert.Nil(db_err.Error)

	assert.Equal(related.ID, user.ID)
}

func TestSessionFromIDDoesntLoadInvalidSession(t *testing.T) {
	assert := assert.New(t)

	session, err := SessionFromID("foopoo")
	assert.NotNil(err)
	assert.Equal("", session.Cookie)
}

// TestHTTPCookie

func TestUserReturnsUserWithSession(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	cookie, err := user.NewSession()
	assert.Nil(err)

	session, err := SessionFromID(cookie)
	assert.Nil(err)

	active, err := session.User()
	assert.Nil(err)

	assert.Equal(active.ID, user.ID)
}
