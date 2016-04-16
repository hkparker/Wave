package database

import (
	"crypto"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
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
	AccountSetupToken  string
	OTPResetNextLogin  bool
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
		Email:             email,
		OTPData:           otp_data,
		AccountSetupToken: helpers.RandomString(),
	}
	db_err := DB().Create(&user)
	if db_err.Error != nil {
		err = db_err.Error
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("error_saving_user")
	} else {
		//user.EmailRegistration()
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"email":  "account_register",
		}).Info("email_sent")
	}
	log.WithFields(log.Fields{
		"UserID": user.ID,
		"email":  user.Email,
	}).Info("user_created")
	return
}

func (user *User) SetPassword(password string) (err error) {
	pw_data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("bcrypt_error")
		return
	}
	user.Password = pw_data
	DB().Save(&user)
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_set")
	return
}

func (user *User) ResetPassword() (err error) {
	user.DestroyAllSessions()
	user.PasswordResetToken = helpers.RandomString()
	db_err := DB().Save(&user)
	if db_err.Error != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
		}).Warn("error_saving_user")
		err = db_err.Error
	} else {
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"email":  "password_reset",
		}).Info("email_sent")
		// email user
	}
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("user_password_reset")
	return
}

func (user *User) ResetTwoFactor() (err error) {
	user.DestroyAllSessions()
	otp, err := twofactor.NewTOTP(user.Email, "Wave", crypto.SHA512, 8)
	if err != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("otp_error")
		return
	}
	otp_data, err := otp.ToBytes()
	if err != nil {
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("otp_error")
		return
	}
	user.OTPData = otp_data
	user.OTPResetNextLogin = true
	db_err := DB().Save(&user)
	if db_err.Error != nil {
		err = db_err.Error
		log.WithFields(log.Fields{
			"UserID": user.ID,
			"error":  err,
		}).Warn("error_saving_user")
		return
	}
	log.WithFields(log.Fields{
		"UserID": user.ID,
	}).Info("two_factor_reset")
	return
}

func ValidAuthentication(email, password, token string) bool {
	// wrap in constant time block
	// lookup user by email in db
	// check password
	// check otp data
	return false
}

func (user *User) NewSession() (wave_session string, err error) {
	now := time.Now()
	wave_session = helpers.RandomString()
	session := Session{
		UserID:            user.ID,
		OriginallyCreated: now,
		LastUsed:          now,
		Cookie:            wave_session,
	}
	user.Sessions = append(user.Sessions, session)
	DB().Save(&user)
	log.WithFields(log.Fields{
		"UserID": user.ID,
		"Time":   now,
		"Cookie": wave_session,
	}).Info("session_created")
	return
}

func (user *User) DestroyAllSessions() {
}
