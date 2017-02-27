package metadata

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func updateKnownDevices(frame models.Wireless80211Frame) {
	if _, ok := Devices[frame.Address1]; !ok {
		registerNewDevice(frame)
	} else if _, ok := Devices[frame.Address2]; !ok {
		registerNewDevice(frame)
	} else if _, ok := Devices[frame.Address3]; !ok {
		registerNewDevice(frame)
	} else if _, ok := Devices[frame.Address4]; !ok {
		registerNewDevice(frame)
	}
}

func registerNewDevice(frame models.Wireless80211Frame) {
	device := models.Device{
		MAC: frame.Address1,
	}
	Devices[frame.Address1] = device
}

func visualizeNewDevice(device models.Device) {
	log.WithFields(log.Fields{
		"mac": device.MAC,
	}).Info("New Device")
	//controllers.VisualPool <-
}
