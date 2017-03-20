package models

import (
	"github.com/jinzhu/gorm"
)

type Network struct {
	gorm.Model
	SSID         string `sql:"not null;unique"`
	Clients      []Device
	AccessPoints []Device
}

func (network *Network) Save() error {
	return Orm.Save(&network).Error
}

func (network *Network) Delete() error {
	return Orm.Delete(&network).Error
}
