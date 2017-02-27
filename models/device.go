package models

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

type Device struct {
	gorm.Model
	MAC         string
	Vendor      string
	AccessPoint bool
}

func (device *Device) Save() error {
	return Orm.Save(&device).Error
}
