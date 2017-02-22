package models

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	Title    string
	Rule     int
	Severity string
}
