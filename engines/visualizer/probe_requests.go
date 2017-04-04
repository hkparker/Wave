package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

// With thanks to:
// https://mrncciew.com/2014/10/27/cwap-802-11-probe-requestresponse/

func updateProbeRequests(frame models.Wireless80211Frame) {
	ssid := string(frame.Elements["SSID"])
	if len(ssid) == 0 {
		var dev models.Device
		ret := models.Orm.Where("MAC = ?", frame.Address2).First(&dev)
		if ret.Error != nil {
			log.WithFields(log.Fields{
				"at":    "visualizer.updateProbeRequests",
				"MAC":   frame.Address2,
				"error": ret.Error,
			}).Error("error looking up Device")
		} else if !dev.Probing {
			dev.Probing = true
			dev.Save()
			Devices[frame.Address2] = dev
			visualizeNullProbe(frame.Address2)
		}
	} else {
		var dev models.Device
		ret := models.Orm.Where("MAC = ?", frame.Address2).First(&dev)
		if ret.Error != nil {
			log.WithFields(log.Fields{
				"at":    "visualizer.updateProbeRequests",
				"MAC":   frame.Address2,
				"error": ret.Error,
			}).Error("error looking up Device")
		} else {
			var network models.Network
			exists := false
			all_networks := make([]models.Network, 0)
			models.Orm.Find(&all_networks)
			for _, net := range all_networks {
				if net.SSID == ssid {
					network = net
					exists = true
				}
			}
			if !exists {
				network = models.Network{
					SSID: ssid,
				}
				network.Save()
				Networks[ssid] = network //append(Networks[ssid], network)
			}

			networks := make([]models.Network, 0)
			models.Orm.Model(&dev).Related(&networks, "ProbedFor")
			associated := false
			for _, net := range networks {
				if net.SSID == ssid {
					associated = true
					break
				}
			}
			if !associated {
				dev.ProbedFor = append(dev.ProbedFor, network)
				dev.Save()
			}

			visualizeProbeRequest(frame.Address2, ssid)
		}
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
