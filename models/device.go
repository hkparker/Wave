package models

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	gorm.Model
	MAC         string `sql:"not null;unique"`
	Vendor      string
	AccessPoint bool
	Probing     bool
	ProbedFor   []Network
	Online      bool
}

func (device *Device) VisualData() map[string]string {
	is_ap := "false"
	if device.AccessPoint {
		is_ap = "true"
	}
	probing := "false"
	if device.Probing {
		probing = "true"
	}
	probed_for := ""
	networks := make([]Network, 0)
	Orm.Model(&device).Related(&networks, "ProbedFor")
	for i, net := range networks {
		probed_for += net.SSID
		if i < len(networks)-1 {
			probed_for += " "
		}
	}
	return map[string]string{
		"type":      "NewDevice",
		"MAC":       device.MAC,
		"Vendor":    device.Vendor,
		"IsAP":      is_ap,
		"Probing":   probing,
		"ProbedFor": probed_for,
	}
}

func (device *Device) Save() error {
	return Orm.Save(&device).Error
}

func (device *Device) Delete() error {
	return Orm.Delete(&device).Error
}
