package models

import (
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCreateUserCreatesUserInCorrectState(t *testing.T) {
	assert := assert.New(t)

	username := helpers.RandomString()
	reset_link, err := CreateUser(username)
	assert.Nil(err)
	assert.NotEqual("", reset_link)

	var user User
	db_err := database.Orm.First(&user, "Username = ?", username)
	assert.Nil(db_err.Error)

	assert.Equal(false, user.Admin)
}

func TestSetPasswordSetsPassword(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	password := "figgindiggle"
	err := user.SetPassword(password)
	assert.Nil(err)

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	assert.Nil(err)
}

func TestResetPasswordResetsPassword(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	user.NewSession()
	user.SetPassword("hunter2")
	assert.Equal("", user.PasswordResetToken)
	err := user.ResetPassword()
	assert.Nil(err)

	assert.Equal(0, len(user.Sessions))
	assert.NotEqual("", user.PasswordResetToken)
	assert.NotEqual(true, user.ValidAuthentication("hunter2"))
}

func TestValidAuthenticationWithValidAuthentication(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	password := "flahblahblah"
	user.SetPassword(password)
	assert.Equal(true, user.ValidAuthentication(password))
}

func TestValidAuthenticationWithBadPassword(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	password := "flahblahblah"
	user.SetPassword(password)
	assert.Equal(false, user.ValidAuthentication("flooblublu"))
}

func TestNewSessionCreatesSession(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	cookie, err := user.NewSession()
	assert.Nil(err)

	var session Session
	db_err := database.Orm.First(&session, "Cookie = ?", cookie)
	assert.Nil(db_err.Error)
	if assert.NotNil(session.UserID) {
		assert.Equal(session.UserID, user.ID)
	}
}

func TestDestroyAllSessionsDestroysAllSessions(t *testing.T) {
	assert := assert.New(t)

	user := CreateTestUser([]string{})
	_, err := user.NewSession()
	assert.Nil(err)

	assert.Equal(1, len(user.Sessions))
	user.DestroyAllSessions()
	assert.Equal(0, len(user.Sessions))
}
