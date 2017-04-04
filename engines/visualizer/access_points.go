package visualizer

import (
	"bytes"
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
	}
	if !dev.AccessPoint {
		dev.AccessPoint = true
		dev.Save()
		Devices[frame.Address3] = dev
	}

	//ssid := string(frame.Elements["SSID"])
	ssid := string(bytes.Trim(frame.Elements["SSID"], "\x00"))
	var net models.Network
	ret = models.Orm.Where("SS_ID = ?", ssid).First(&net)
	if ret.Error != nil {
		net = models.Network{
			AccessPoints: make([]models.Device, 0),
			SSID:         ssid,
		}
		models.Orm.Create(&net)
		Networks[ssid] = net
	}

	var aps []models.Device
	models.Orm.Model(&net).Related(&aps, "AccessPoints").Where("MAC = ?", frame.Address3)
	if len(aps) == 0 {
		net.AccessPoints = append(net.AccessPoints, dev)
		net.Save()
		Networks[ssid] = net
		visualizeNewAP(frame)
	}
}

func visualizeNewAP(frame models.Wireless80211Frame) {
	ssid := string(frame.Elements["SSID"])
	VisualEvents <- VisualEvent{
		TYPE:        TYPE_UPDATE_AP,
		DEVICE_MAC:  frame.Address3,
		DEVICE_ISAP: "true",
		SSID:        ssid,
	}
	log.WithFields(log.Fields{
		"at":   "visualizer.visualizeNewAP",
		"mac":  frame.Address3,
		"ssid": ssid,
	}).Debug("update device as ap")
}
