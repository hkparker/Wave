package visualizer

import (
	log "github.com/sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
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
	ret := true
	association1 := Associations[mac1]
	if !helpers.StringIncludedIn(association1, mac2) {
		Associations[mac1] = append(Associations[mac1], mac2)
		ret = false
	}
	association2 := Associations[mac2]
	if !helpers.StringIncludedIn(association2, mac1) {
		Associations[mac2] = append(Associations[mac2], mac1)
		ret = false
	}
	//if mac1 == mac2 {
	//	log.Warn("association between identical mac " + mac1)
	//	return true
	//}
	return ret
}

func visualizeAssociation(mac1, mac2 string) {
	VisualEvents <- VisualEvent{
		TYPE:   "NewAssociation",
		"MAC1": mac1,
		"MAC2": mac2,
	}
	log.WithFields(log.Fields{
		"at":   "visualizer.visualizeAssociation",
		"mac1": mac1,
		"mac2": mac2,
	}).Debug("associate devices")
}
