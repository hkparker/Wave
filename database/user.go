package database

import (
	"crypto"
	"crypto/rand"
	"github.com/jbenet/go-base58"
	"github.com/jinzhu/gorm"
	"github.com/sec51/twofactor"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Name               string
	Password           []byte
	Email              string `sql:"not null;unique"`
	Admin              bool
	Sessions           []Session
	OTPData            []byte
	PasswordResetToken string
}

func CreateUser(email string) (err error) {
	otp, err := twofactor.NewTOTP(email, "Wave", crypto.SHA512, 8)
	if err != nil {
		return
	}
	otp_data, err := otp.ToBytes()
	if err != nil {
		return
	}
	user := User{
		Email:              email,
		OTPData:            otp_data,
		PasswordResetToken: "",
	}
	DB().Create(&user)
	// email the user the register link
	return
}

func (user *User) ResetPassword() {
	user.DestroyAllSessions()
	// setup one time password page
	// email user
}

func (user *User) SetPassword(password string) (err error) {
	pw_data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = pw_data
	DB().Save(&user)
	return
}

func (user *User) ResetTwoFactor() {
	otp, err := twofactor.NewTOTP(user.Email, "Wave", crypto.SHA512, 8)
	if err != nil {
		return
	}
	otp_data, err := otp.ToBytes()
	if err != nil {
		return
	}
	user.OTPData = otp_data
	// create a database record for impending reset link
	DB().Save(&user)
	// email the user
}

func ValidAuthentication(email, password, token string) bool {
	// wrap in constant time block
	// lookup user by email in db
	// check password
	// check otp data
	return false
}

func (user *User) NewSession() (wave_session string, err error) {
	session_bytes := make([]byte, 32)
	_, err = rand.Read(session_bytes)
	if err != nil {
		return
	}
	wave_session = base58.Encode(session_bytes)
	now := time.Now()
	session := Session{
		UserID:            user.ID,
		OriginallyCreated: now,
		LastUsed:          now,
		Cookie:            wave_session,
	}
	user.Sessions = append(user.Sessions, session)
	DB().Save(&user)
	return
}

func (user *User) DestroyAllSessions() {
}
