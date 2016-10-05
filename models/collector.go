package models

import (
	//log "github.com/Sirupsen/logrus"
	//"github.com/gin-gonic/gin"
	//"github.com/hkparker/Wave/database"
	//"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
	//"time"
)

type Collector struct {
	gorm.Model
	Name       string
	ClientCert string
}
