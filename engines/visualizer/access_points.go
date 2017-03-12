package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func updateAccessPoints(frame models.Wireless80211Frame) {
	DevicesMux.Lock()
	defer DevicesMux.Unlock()
	if frame.Type == "MgmtBeacon" {
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
	}
}

func visualizeNewAP(mac string) {
	update_resources := make(VisualEvent)
	update_resources["UpdateDevices"] = append(
		update_resources["UpdateDevices"],
		map[string]string{
			"MAC":  mac,
			"IsAP": "true",
		},
	)
	VisualEvents <- update_resources
	log.WithFields(log.Fields{
		"at":  "visualizer.visualizeNewAP",
		"mac": mac,
	}).Info("update device as ap")
}
