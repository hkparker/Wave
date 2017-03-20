package visualizer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	"strings"
	"sync"
)

const (
	NEW_DEVICES         = "NewDevices"
	UPDATE_DEVICES      = "UpdateDevices"
	DRAW_EVENTS         = "DrawEvents"
	NEW_ASSOCIATIONS    = "NewAssociations"
	UPDATE_ASSOCIATIONS = "UpdateAssociationS"

	DEVICE_ISAP           = "IsAP"
	DEVICE_MAC            = "MAC"
	DEVICE_NULLPROBE      = "NullProbe"
	DEVICE_PROBE          = "ProbedFor"
	DEVICE_POWERSTATE     = "PowerState"
	DEVICE_POWERSTATE_ON  = "online"
	DEVICE_POWERSTATE_OFF = "offline"

	EVENT_NULL_PROBE    = "NullProbe"
	EVENT_PROBE_REQUEST = "ProbeRequest"

	EVENT = "Event"
	SSID  = "SSID"
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
	DevicesMux.Lock()
	defer DevicesMux.Unlock()
	NetworksMux.Lock()
	defer NetworksMux.Unlock()
	updateKnownDevices(frame)

	if len(frame.Type) < 4 {
		log.WithFields(log.Fields{}).Warn()
		return
	}
	switch frame.Type[:4] {
	case "Mgmt":
		insertMgmt(frame)
	case "Data":
		insertData(frame)
	case "Ctrl":
		insertCtrl(frame)
	//case "Rese":
	default:
		log.WithFields(log.Fields{
			"at":         "visualizer.Insert",
			"frame.Type": frame.Type,
		}).Warn("unknown frame type")
	}
}

func insertMgmt(frame models.Wireless80211Frame) {
	switch frame.Type {
	case "MgmtAssociationReq":
	case "MgmtAssociationResp":
	case "MgmtReassociationReq":
	case "MgmtReassociationResp":
	case "MgmtProbeReq":
		updateProbeRequests(frame)
	case "MgmtProbeResp":
	case "MgmtMeasurementPilot":
	case "MgmtBeacon":
		updateAccessPoints(frame)
	case "MgmtATIM":
	case "MgmtDisassociation":
	case "MgmtAuthentication":
	case "MgmtDeauthentication":
	case "MgmtAction":
	case "MgmtActionNoAck":
	default:
		log.WithFields(log.Fields{
			"at":   "visualizer.insertMgmt",
			"type": frame.Type,
		}).Warn("unknown frame type")
	}
}

func insertData(frame models.Wireless80211Frame) {
	//updateNetworkAssociations(frames)
	//updateTx()
	switch frame.Type {
	case "Data":
	case "DataCFAck":
	case "DataCFPoll":
	case "DataCFAckPoll":
	case "DataNull":
		updateDataNull(frame)
	case "DataCFAckNoData":
	case "DataCFPollNoData":
	case "DataCFAckPollNoData":
	case "DataQOSData":
	case "DataQOSDataCFAck":
	case "DataQOSDataCFPoll":
	case "DataQOSDataCFAckPoll":
	case "DataQOSNull":
	case "DataQOSCFPollNoData":
	case "DataQOSCFAckPollNoData":
	default:
		log.WithFields(log.Fields{
			"at":   "visualizer.insertData",
			"type": frame.Type,
		}).Warn("unknown frame type")
	}
}

func insertCtrl(frame models.Wireless80211Frame) {
	switch frame.Type {
	case "CtrlWrapper":
	case "CtrlBlockAckReq":
	case "CtrlBlockAck":
	case "CtrlPowersavePoll":
	case "CtrlRTS":
	case "CtrlCTS":
	case "CtrlAck":
	case "CtrlCFEnd":
	case "CtrlCFEndAck":
	default:
		log.WithFields(log.Fields{
			"at":   "visualizer.insertCtrl",
			"type": frame.Type,
		}).Warn("unknown frame type")
	}
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
