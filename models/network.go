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

func (network *Network) VisualData() []map[string]string {
	set := make([]map[string]string, 0)
	aps := make([]Device, 0)
	Orm.Model(&network).Related(&aps, "AccessPoints")
	for _, ap := range aps {
		set = append(set, map[string]string{
			"type": "UpdateAccessPoint",
			"MAC":  ap.MAC,
			"IsAP": "true",
			"SSID": network.SSID,
		})
	}
	return set
}

func (network *Network) Save() error {
	return Orm.Save(&network).Error
}

func (network *Network) Delete() error {
	return Orm.Delete(&network).Error
}
