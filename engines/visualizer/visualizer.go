package visualizer

import (
	"github.com/hkparker/Wave/models"
	"sync"
)

type VisualEvent map[string][]map[string]string

var VisualEvents = make(chan VisualEvent, 0)
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

func CatchupEvents() []VisualEvent {
	new_resources := make(VisualEvent)
	for _, device := range Devices {
		new_resources["NewDevices"] = append(
			new_resources["NewDevices"],
			device.VisualData(),
		)
	}
	// add other resources, create other events
	return []VisualEvent{
		new_resources,
	}
}
