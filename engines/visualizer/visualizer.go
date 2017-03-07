package visualizer

import (
	"github.com/hkparker/Wave/models"
	"sync"
)

type VisualEvent map[string][]map[string]string

var VisualEvents = make(chan interface{}, 0)
var Devices = make(map[string]models.Device)
var DevicesMux sync.Mutex
var Networks = make(map[string][]models.Network)
var NetworksMux sync.Mutex

func Load() {
	DevicesMux.Lock()
	defer DevicesMux.Unlock()
	NetworksMux.Lock()
	defer NetworksMux.Unlock()

	all_devices := make([]models.Device, 0)
	models.Orm.Find(&all_devices)
	for _, device := range all_devices {
		Devices[device.MAC] = device
	}

	all_networks := make([]models.Network, 0)
	models.Orm.Find(&all_networks)
	for _, network := range all_networks {
		Networks[network.SSID] = append(Networks[network.SSID], network)
	}
}

func Insert(frame models.Wireless80211Frame) {
	updateKnownDevices(frame)
	updateAccessPoints(frame)
	//updateProbeRequests(frame)
	//updateNetworkAssociations(frames)
	//updateTx()
}

func CatchupEvents() []map[string]string {
	// for each device, visualize new device
	// for each session, ...
	return make([]map[string]string, 0)
}
