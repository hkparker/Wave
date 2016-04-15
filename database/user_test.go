package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserCreatesUserInCorrectState(t *testing.T) {

}

func TestNewSessionCreatesSession(t *testing.T) {
	assert := assert.New(t)

	user := TestUser([]string{})
	cookie, err := user.NewSession()
	assert.Nil(err)
	session := Session{}
	db_err := DB().First(&session, "Cookie = ?", cookie)
	assert.Nil(db_err.Error)
	if assert.NotNil(session.UserID) {
		assert.Equal(session.UserID, user.ID)
	}
}
