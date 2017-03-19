package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
	"strings"
)

func updateKnownDevices(frame models.Wireless80211Frame) {
	if _, ok := Devices[frame.Address1]; !ok {
		DevicesMux.Lock()
		if len(frame.Address1) == 0 {
			return
		}
		registerNewDevice(frame.Address1)
		DevicesMux.Unlock()
	}
	if _, ok := Devices[frame.Address2]; !ok {
		DevicesMux.Lock()
		if len(frame.Address2) == 0 {
			return
		}
		registerNewDevice(frame.Address2)
		DevicesMux.Unlock()
	}
	if _, ok := Devices[frame.Address3]; !ok {
		DevicesMux.Lock()
		if len(frame.Address3) == 0 {
			return
		}
		registerNewDevice(frame.Address3)
		DevicesMux.Unlock()
	}
	if _, ok := Devices[frame.Address4]; !ok {
		DevicesMux.Lock()
		if len(frame.Address4) == 0 {
			return
		}
		registerNewDevice(frame.Address4)
		DevicesMux.Unlock()
	}
}

func registerNewDevice(mac string) {
	if len(mac) != 17 {
		log.WithFields(log.Fields{
			"at":  "visualizer.registerDevice",
			"mac": mac,
		}).Warn("malformed mac")
		return
	}
	if broadcast(mac) {
		return
	}
	vendor := "unknown"
	if len(mac) >= 8 {
		if vendor_string, ok := VendorBytes[strings.ToUpper(mac[0:8])]; ok {
			vendor = vendor_string
		}
	}
	device := models.Device{
		MAC:    mac,
		Vendor: vendor,
	}
	Devices[mac] = device
	device.Save()
	visualizeNewDevice(device)
}

func broadcast(mac string) bool {
	if mac == "ff:ff:ff:ff:ff:ff" {
		return true
	}
	return false
}

func visualizeNewDevice(device models.Device) {
	new_resources := make(VisualEvent)
	new_resources["NewDevices"] = append(
		new_resources["NewDevices"],
		device.VisualData(),
	)
	VisualEvents <- new_resources
	log.WithFields(log.Fields{
		"at":  "visualizeNewDevice",
		"mac": device.MAC,
	}).Debug("new device observed")
}
