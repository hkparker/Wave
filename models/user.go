package models

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	"github.com/sec51/twofactor"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Name            string
	Password        []byte
	Email           string `sql:"not null;unique"`
	Sessions        []Session
	OTPData         []byte
	OTPReset        bool
	PasswordResetID string
}

func RegisterUser(email string) (err error) {
	// if the email already exists, return an error
	otp, err := twofactor.NewTOTP(email, "Wave", crypto.SHA512, 8)
	if err != nil {
		return
	}
	otp_data, err := otp.ToBytes()
	if err != nil {
		return
	}
	_ = User{
		Email:           email,
		OTPData:         otp_data,
		OTPReset:        true,
		PasswordResetID: "",
	}
	//db.Create(&user)
	// email the user the register link
	return
}

func (user *User) ResetPassword() {
	// invalidate all sessions
	// setup one time password page
	// email user
}

func (user *User) SetPassword(password string) (err error) {
	pw_data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = pw_data
	//db.Save(&user)
	return
}

func (user *User) ResetAuthentication() {
	// user will need to setup new two factor code next time it is asked for
}

func (user *User) ValidAuthentication(email, password, token string) bool {
	return false
}

func (user *User) NewSession() (wave_session string, err error) {
	session_bytes := make([]byte, 32)
	_, err = rand.Read(session_bytes)
	if err != nil {
		return
	}
	wave_session = base64.StdEncoding.EncodeToString(session_bytes)
	session := Session{
		LastUsed: time.Now(),
		Cookie:   wave_session,
	}
	user.Sessions = append(user.Sessions, session)
	helpers.DB().Save(&user)
	return
}

func (user *User) DestroySession(cookie string) {

}

func (user *User) DestroyAllSessions() {
}
