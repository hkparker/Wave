package visualizer

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func updateAccessPoints(frame models.Wireless80211Frame) {
	if frame.Type == "MgmtBeacon" {
		// Mgmt frame BSSID is Address3
		dev := Devices[frame.Address3]
		if !dev.AccessPoint {
			dev.AccessPoint = true
			dev.Save()
			Devices[frame.Address3] = dev
			visualizeNewAP(frame.Address3)
		}
	}
}

func visualizeNewAP(mac string) {
	//controllers.VisualPool <-
	//log.WithFields(log.Fields{
	//	"mac": mac,
	//}).Info("new AP observed")
}
