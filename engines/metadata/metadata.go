package metadata

import (
	"github.com/hkparker/Wave/models"
)

var Devices = make(map[string]models.Device)

func init() {
	// Load devices from DB
}

func Insert(frame models.Wireless80211Frame) {
	updateKnownDevices(frame)
	updateAccessPoints(frame)
	//updateProbeRequests(frame)
	//updateNetworkAssociations(frames)
	//updateTxRx()
}
