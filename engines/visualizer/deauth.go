package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
)

// With thanks to:
// https://mrncciew.com/2014/10/11/802-11-mgmt-deauth-disassociation-frames/

func animateDeauth(frame models.Wireless80211Frame) {
	if frame.Address3 == frame.Address1 {
		log.Warn("duplicate deauth mac " + frame.Address1)
		return
	}
	if links, ok := Associations[frame.Address1]; ok {
		if helpers.StringIncludedIn(links, frame.Address3) {
			Associations[frame.Address1] = helpers.StringsExcept(links, frame.Address3)
		}
	}
	if links, ok := Associations[frame.Address3]; ok {
		if helpers.StringIncludedIn(links, frame.Address1) {
			Associations[frame.Address3] = helpers.StringsExcept(links, frame.Address1)
		}
	}

	VisualEvents <- VisualEvent{
		"type":   "AnimateDeauth",
		"Target": frame.Address1,
		"Source": frame.Address3,
	}
	log.WithFields(log.Fields{
		"at": "animateDeauth",
	}).Debug("animating deauth")
}
