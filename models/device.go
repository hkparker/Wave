package models

import (
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

func (device *Device) Delete() error {
	return Orm.Delete(&device).Error
}
