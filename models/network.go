package models

import (
	"github.com/jinzhu/gorm"
)

type Network struct {
	gorm.Model
	SSID         string
	Clients      []Device
	AccessPoints []Device
}

func (network *Network) Save() error {
	return Orm.Save(&network).Error
}

func (network *Network) Delete() error {
	return Orm.Delete(&network).Error
}
