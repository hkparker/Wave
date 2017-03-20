package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

// With thanks to:
// https://mrncciew.com/2014/10/27/cwap-802-11-probe-requestresponse/

func updateProbeRequests(frame models.Wireless80211Frame) {
	if len(frame.Elements["SSID"]) == 0 {
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
		} else if !dev.Probing {
			// add the network to the ProbedFor association
			//visualizeProbeRequest(frame)
		}
	}
}

func visualizeNullProbe(mac string) {
	update_resources := make(VisualEvent)
	update_resources[UPDATE_DEVICES] = append(
		update_resources[UPDATE_DEVICES],
		map[string]string{
			DEVICE_MAC:       mac,
			DEVICE_NULLPROBE: "",
		},
	)
	VisualEvents <- update_resources
	log.WithFields(log.Fields{
		"at":  "visualizer.visualizeNullProbe",
		"mac": mac,
	}).Debug("visualizing device null probe")
}
