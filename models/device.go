package models

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	gorm.Model
	MAC         string `sql:"not null;unique"`
	Vendor      string
	AccessPoint bool
}

func (device *Device) VisualData() map[string]string {
	return map[string]string{
		"MAC":    device.MAC,
		"Vendor": device.Vendor,
	}
}

func (device *Device) Save() error {
	return Orm.Save(&device).Error
}

func (device *Device) Delete() error {
	return Orm.Delete(&device).Error
}
