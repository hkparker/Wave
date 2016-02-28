package models

import (
	"crypto"
	"github.com/jinzhu/gorm"
	"github.com/sec51/twofactor"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name            string
	Password        []byte
	Email           string
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

func (user *User) ResetTwoFactor() {
	// user will need to setup new two factor code next time it is asked for
}

func (user *User) Login() {

}
