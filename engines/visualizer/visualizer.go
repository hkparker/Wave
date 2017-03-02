package visualizer

import (
	"github.com/hkparker/Wave/models"
	"sync"
)

var Devices = make(map[string]models.Device)
var DevicesMux sync.Mutex

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

func CatchupEvents() []map[string]string {
	// for each device, visualize new device
	// for each session, ...
	return make([]map[string]string, 0)
}
