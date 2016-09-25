package models

import (
	"github.com/hkparker/Wave/database"
	//"github.com/hkparker/Wave/helpers"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func init() {
	test = true
}

func TestCreateUserCreatesUserInCorrectState(t *testing.T) {
	//assert := assert.New(t)

	//email := helpers.RandomString() + "@example.com"
	//err := CreateUser(email)
	//assert.Nil(err)

	//var user User
	//db_err := database.Orm.First(&user, "Email = ?", email)
	//assert.Nil(db_err.Error)

	//assert.Equal(false, user.Admin)
}

func TestCurrentUserGetsUserForSession(t *testing.T) {

}

func TestSetPasswordSetsPassword(t *testing.T) {
	assert := assert.New(t)

	user := TestUser([]string{})
	password := "figgindiggle"
	err := user.SetPassword(password)
	assert.Nil(err)

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	assert.Nil(err)
}

func TestResetPasswordResetsPassword(t *testing.T) {
	assert := assert.New(t)

	user := TestUser([]string{})
	user.NewSession()
	assert.Equal("", user.PasswordResetToken)
	err := user.ResetPassword()
	assert.Nil(err)

	assert.Equal(0, len(user.Sessions))
	assert.NotEqual("", user.PasswordResetToken)
}

func TestValidAuthenticationWithValidAuthentication(t *testing.T) {
	//assert := assert.New(t)

	//user := TestUser([]string{})
	//password := "flahblahblah"
	//user.SetPassword(password)
	//assert.Equal(true, ValidAuthentication(user.Email, password))
}

func TestValidAuthenticationWithBadEmail(t *testing.T) {
	//assert := assert.New(t)

	//user := TestUser([]string{})
	//password := "flahblahblah"
	//user.SetPassword(password)
	//assert.Equal(false, ValidAuthentication("thisdoesntexist@example.com", password))
}

func TestValidAuthenticationWithBadPassword(t *testing.T) {
	//assert := assert.New(t)

	//user := TestUser([]string{})
	//password := "flahblahblah"
	//user.SetPassword(password)
	//assert.Equal(false, ValidAuthentication(user.Email, "flooblublu"))
}

func TestNewSessionCreatesSession(t *testing.T) {
	assert := assert.New(t)

	user := TestUser([]string{})
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

	user := TestUser([]string{})
	_, err := user.NewSession()
	assert.Nil(err)

	assert.Equal(1, len(user.Sessions))
	//user.DestroyAllSessions()
	//assert.Equal(0, len(user.Sessions))
}
