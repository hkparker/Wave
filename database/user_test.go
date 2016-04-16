package database

import (
	"github.com/hkparker/Wave/helpers"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCreateUserCreatesUserInCorrectState(t *testing.T) {
	assert := assert.New(t)

	email := helpers.RandomString() + "@example.com"
	err := CreateUser(email)
	assert.Nil(err)

	var user User
	db_err := DB().First(&user, "Email = ?", email)
	assert.Nil(db_err.Error)

	assert.NotEqual("", user.AccountSetupToken)
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

func TestResetTwoFactorResetsTwoFactor(t *testing.T) {
	assert := assert.New(t)

	user := TestUser([]string{})
	user.NewSession()
	otp_data := user.OTPData
	err := user.ResetTwoFactor()
	assert.Nil(err)

	assert.Equal(0, len(user.Sessions))
	assert.NotEqual(user.OTPData, otp_data)
	assert.Equal(true, user.OTPResetNextLogin)
}

func TestValidAuthenticationWithValidAuthentication(t *testing.T) {

}

func TestValidAuthenticationWithBadEmail(t *testing.T) {

}

func TestValidAuthenticationWithBadPassword(t *testing.T) {

}

func TestValidAuthenticationWithBadTwoFactor(t *testing.T) {

}

func TestNewSessionCreatesSession(t *testing.T) {
	assert := assert.New(t)

	user := TestUser([]string{})
	cookie, err := user.NewSession()
	assert.Nil(err)

	var session Session
	db_err := DB().First(&session, "Cookie = ?", cookie)
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
	user.DestroyAllSessions()
	assert.Equal(0, len(user.Sessions))
}
