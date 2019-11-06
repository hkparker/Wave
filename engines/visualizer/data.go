package visualizer

import (
	"github.com/hkparker/Wave/models"
	log "github.com/sirupsen/logrus"
)

func updateAssociation(frame models.Wireless80211Frame) {
	mac1 := frame.Address1
	mac2 := frame.Address3

	if frame.Flags80211.ToDS() {
		if frame.Flags80211.FromDS() {
			//log.Warn("11")
		} else {
			//log.Warn("10")
		}
	} else {
		if frame.Flags80211.FromDS() {
			//log.Warn("01")
		} else {
			//log.Warn("00")
		}
	}

	if !memoized(mac1, mac2) {
		visualizeAssociation(mac1, mac2)
	}
}

func memoized(mac1, mac2 string) bool {
	if mac1 == "" || mac2 == "" {
		log.Warn("null mac")
		return true
	}
	if broadcast(mac1) || broadcast(mac2) {
		return true
	}

	key := deterministicKey(mac1, mac2)
	if _, exists := Associations[key]; exists {
		return true
	} else {
		Associations[key] = models.Association{
			Source: mac1,
			Target: mac2,
		}
	}
	return false

	//if mac1 == mac2 {
	//	log.Warn("association between identical mac " + mac1)
	//	return true
	//}
}

func visualizeAssociation(mac1, mac2 string) {
	VisualEvents <- VisualEvent{
		TYPE:   "NewAssociation",
                "Key":  deterministicKey(mac1, mac2),
		"target": mac1,
		"source": mac2,
	}
	log.WithFields(log.Fields{
		"at":   "visualizer.visualizeAssociation",
		"mac1": mac1,
		"mac2": mac2,
	}).Debug("associate devices")
}
