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
	is_ap := "false"
	if device.AccessPoint {
		is_ap = "true"
	}
	return map[string]string{
		"MAC":    device.MAC,
		"Vendor": device.Vendor,
		"IsAP":   is_ap,
	}
}

func (device *Device) Save() error {
	return Orm.Save(&device).Error
}

func (device *Device) Delete() error {
	return Orm.Delete(&device).Error
}
