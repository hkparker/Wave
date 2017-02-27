package models

import (
	//log "github.com/Sirupsen/logrus"
	//"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
)

type Rule struct {
	gorm.Model
	Name     string
	Function string
}
