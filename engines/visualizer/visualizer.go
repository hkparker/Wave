package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	"strings"
	"sync"
)

type VisualEvent map[string][]map[string]string

var VisualEvents = make(chan VisualEvent, 0)
var Devices = make(map[string]models.Device)
var DevicesMux sync.Mutex
var Networks = make(map[string][]models.Network)
var NetworksMux sync.Mutex
var VendorBytes = make(map[string]string)

func init() {
	loadMetadata()
}

func loadMetadata() {
	prefix_path := "engines/visualizer/metadata/nmap-mac-prefixes"
	vendor_data, err := helpers.Asset(prefix_path)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "visualizer.loadMetadata",
			"error": err.Error(),
		}).Error("unable to load vendor bytes")
		return
	}
	lines := strings.Split(string(vendor_data), "\n")
	for _, line := range lines {
		if len(line) == 0 || string(line[0]) == "#" {
			continue
		}
		raw_mac := line[0:6]
		name := line[7:]
		mac := raw_mac[0:2] + ":" +
			raw_mac[2:4] + ":" +
			raw_mac[4:6]
		VendorBytes[strings.ToUpper(mac)] = name
	}
}

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
