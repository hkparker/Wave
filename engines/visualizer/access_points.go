package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func updateAccessPoints(frame models.Wireless80211Frame) {
	// Mgmt frame BSSID is Address3
	var dev models.Device
	ret := models.Orm.Where("MAC = ?", frame.Address3).First(&dev)
	if ret.Error != nil {
		log.WithFields(log.Fields{
			"at":    "visualizer.updateAccessPoints",
			"MAC":   frame.Address3,
			"error": ret.Error,
		}).Error("error looking up AP")
	} else if !dev.AccessPoint {
		dev.AccessPoint = true
		dev.Save()
		Devices[frame.Address3] = dev
		visualizeNewAP(frame.Address3)
	}
	// parse ssid
	// lookup network or create with ssid
	// add device as ap
}

func visualizeNewAP(mac string) {
	VisualEvents <- VisualEvent{
		TYPE:        TYPE_UPDATE_AP,
		DEVICE_MAC:  mac,
		DEVICE_ISAP: "true",
		// SSID
	}
	log.WithFields(log.Fields{
		"at":  "visualizer.visualizeNewAP",
		"mac": mac,
	}).Debug("update device as ap")
}
