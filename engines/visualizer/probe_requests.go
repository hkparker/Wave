package visualizer

import (
	log "github.com/sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

// With thanks to:
// https://mrncciew.com/2014/10/27/cwap-802-11-probe-requestresponse/

func updateProbeRequests(frame models.Wireless80211Frame) {
	ssid := string(frame.Elements["SSID"])
	dev := Devices[frame.Address2]
	if len(ssid) == 0 {
		if !dev.Probing {
			dev.Probing = true
			Devices[frame.Address2] = dev
			visualizeNullProbe(frame.Address2)
		}
	} else {
		exists := false
		var network models.Network
		for _, net := range Networks {
			if net.SSID == ssid {
				network = net
				exists = true
			}
		}
		if !exists {
			network = models.Network{
				SSID: ssid,
			}
			Networks[ssid] = network
		}

		associated := false
		for _, net := range dev.ProbedFor {
			if net.SSID == ssid {
				associated = true
				break
			}
		}
		if !associated {
			dev.ProbedFor = append(dev.ProbedFor, network)
		}

		visualizeProbeRequest(frame.Address2, ssid)
	}
}

func visualizeNullProbe(mac string) {
	VisualEvents <- VisualEvent{
		TYPE:       TYPE_NULL_PROBE_REQUEST,
		DEVICE_MAC: mac,
	}
	log.WithFields(log.Fields{
		"at":       "visualizer.visualizeNullProbe",
		DEVICE_MAC: mac,
	}).Debug("visualizing null probe")
}

func visualizeProbeRequest(mac string, ssid string) {
	VisualEvents <- VisualEvent{
		TYPE:       TYPE_PROBE_REQUEST,
		SSID:       ssid,
		DEVICE_MAC: mac,
	}
	log.WithFields(log.Fields{
		"at":       "visualizer.visualizeProbeRequest",
		SSID:       ssid,
		DEVICE_MAC: mac,
	}).Debug("visualizing probe request")
}
