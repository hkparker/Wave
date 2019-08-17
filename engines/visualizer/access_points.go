package visualizer

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func updateAccessPoints(frame models.Wireless80211Frame) {
	// Mgmt frame BSSID is Address3
	dev := Devices[frame.Address3]
	if !dev.AccessPoint {
		dev.AccessPoint = true
		Devices[frame.Address3] = dev
	}

	//ssid := string(frame.Elements["SSID"])
	ssid := string(bytes.Trim(frame.Elements["SSID"], "\x00"))
	if ssid == "" {
		// Skip empty beacons (Closed System)
		return
	}

	net, ok := Networks[ssid]
	if !ok {
		net = models.Network{
			SSID: ssid,
		}
		Networks[ssid] = net
	}

	associated := false
	for _, ap := range net.AccessPoints {
		if ap.MAC == dev.MAC {
			associated = true
			break
		}
	}
	if !associated {
		net.AccessPoints = append(net.AccessPoints, dev)
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
