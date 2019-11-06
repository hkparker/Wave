package visualizer

import (
	"github.com/hkparker/Wave/models"
	log "github.com/sirupsen/logrus"
)

// With thanks to:
// https://mrncciew.com/2014/10/11/802-11-mgmt-deauth-disassociation-frames/

func animateDeauth(frame models.Wireless80211Frame) {
	if frame.Address3 == frame.Address1 {
		log.Warn("duplicate deauth mac " + frame.Address1)
		return
	}
	key := deterministicKey(frame.Address1, frame.Address3)
	_, ok := Associations[key]
	if !ok {
		// deauth on unknown association
		return
	}
	delete(Associations, key)

	VisualEvents <- VisualEvent{
		"type":   "AnimateDeauth",
		"Target": frame.Address1,
		"Source": frame.Address3,
	}
	log.WithFields(log.Fields{
		"at": "animateDeauth",
	}).Debug("animating deauth")
}
