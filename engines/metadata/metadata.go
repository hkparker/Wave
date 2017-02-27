package metadata

import (
	"github.com/hkparker/Wave/models"
)

var Devices = make(map[string]models.Device)

func Insert(frame models.Wireless80211Frame) {
	updateKnownDevices(frame)
	updateProbeRequests(frame)
	updateAccessPoints(frame)
	//updateNetworkAssociations(frames)
	//updateTxRx()
}
