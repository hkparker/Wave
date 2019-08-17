package visualizer

import (
	log "github.com/sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	"strings"
)

func updateKnownDevices(frame models.Wireless80211Frame) {
	if _, ok := Devices[frame.Address1]; !ok {
		if len(frame.Address1) == 0 {
			return
		}
		registerNewDevice(frame.Address1)
	}
	if _, ok := Devices[frame.Address2]; !ok {
		if len(frame.Address2) == 0 {
			return
		}
		registerNewDevice(frame.Address2)
	}
	if _, ok := Devices[frame.Address3]; !ok {
		if len(frame.Address3) == 0 {
			return
		}
		registerNewDevice(frame.Address3)
	}
	if _, ok := Devices[frame.Address4]; !ok {
		if len(frame.Address4) == 0 {
			return
		}
		registerNewDevice(frame.Address4)
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
		} else {
			//log.Warn("unknown vendor for " + mac)
			//return
		}
	}
	device := models.Device{
		MAC:    mac,
		Vendor: vendor,
	}
	Devices[mac] = device
	visualizeNewDevice(device)
}

func broadcast(mac string) bool {
	// IPv6 multicast DHCP
	if mac[:5] == "33:33" {
		return true
	}
	// IPv4 multicast DHCP
	if mac[:8] == "01:00:5e" {
		return true
	}
	return helpers.StringIncludedIn(
		[]string{
			"ff:ff:ff:ff:ff:ff",
			"33:33:00:01:00:02",
			"00:00:00:00:00:00",
		},
		mac,
	)
}

func visualizeNewDevice(device models.Device) {
	VisualEvents <- device.VisualData()
	log.WithFields(log.Fields{
		"at":  "visualizeNewDevice",
		"mac": device.MAC,
	}).Debug("new device observed")
}
