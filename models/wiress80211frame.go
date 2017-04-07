package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// https://github.com/torvalds/linux/blob/master/include/linux/ieee80211.h#L1787
var ELEMENT_IDS = map[byte]string{
	0:   "SSID",
	1:   "SUPPORTED_RATES",
	5:   "TRAFFIC_INDICATION_MAP",
	7:   "COUNTRY",
	11:  "QBSS",
	42:  "ERP_INFO",
	48:  "RSN",
	50:  "EXTENDED_SUPPORTED_RATES",
	61:  "HT_OPERATION",
	3:   "DS_PARAMETER_SET",
	45:  "HT_CAPABILITIES",
	127: "EXTENDED_CAPABILITIES",
	221: "VENDOR_SPECIFIC",
	/*
	   WLAN_EID_FH_PARAMS = 2, // reserved now
	   WLAN_EID_DS_PARAMS = 3,
	   WLAN_EID_CF_PARAMS = 4,
	   WLAN_EID_TIM = 5,
	   WLAN_EID_IBSS_PARAMS = 6,
	   WLAN_EID_COUNTRY = 7,
	   // 8, 9 reserved
	   WLAN_EID_REQUEST = 10,
	   WLAN_EID_QBSS_LOAD = 11,
	   WLAN_EID_EDCA_PARAM_SET = 12,
	   WLAN_EID_TSPEC = 13,
	   WLAN_EID_TCLAS = 14,
	   WLAN_EID_SCHEDULE = 15,
	   WLAN_EID_CHALLENGE = 16,
	   // 17-31 reserved for challenge text extension
	   WLAN_EID_PWR_CONSTRAINT = 32,
	   WLAN_EID_PWR_CAPABILITY = 33,
	   WLAN_EID_TPC_REQUEST = 34,
	   WLAN_EID_TPC_REPORT = 35,
	   WLAN_EID_SUPPORTED_CHANNELS = 36,
	   WLAN_EID_CHANNEL_SWITCH = 37,
	   WLAN_EID_MEASURE_REQUEST = 38,
	   WLAN_EID_MEASURE_REPORT = 39,
	   WLAN_EID_QUIET = 40,
	   WLAN_EID_IBSS_DFS = 41,

	   WLAN_EID_TS_DELAY = 43,
	   WLAN_EID_TCLAS_PROCESSING = 44,
	   WLAN_EID_HT_CAPABILITY = 45,
	   WLAN_EID_QOS_CAPA = 46,
	   // 47 reserved for Broadcom

	   WLAN_EID_802_15_COEX = 49,
	   WLAN_EID_EXT_SUPP_RATES = 50,
	   WLAN_EID_AP_CHAN_REPORT = 51,
	   WLAN_EID_NEIGHBOR_REPORT = 52,
	   WLAN_EID_RCPI = 53,
	   WLAN_EID_MOBILITY_DOMAIN = 54,
	   WLAN_EID_FAST_BSS_TRANSITION = 55,
	   WLAN_EID_TIMEOUT_INTERVAL = 56,
	   WLAN_EID_RIC_DATA = 57,
	   WLAN_EID_DSE_REGISTERED_LOCATION = 58,
	   WLAN_EID_SUPPORTED_REGULATORY_CLASSES = 59,
	   WLAN_EID_EXT_CHANSWITCH_ANN = 60,

	   WLAN_EID_SECONDARY_CHANNEL_OFFSET = 62,
	   WLAN_EID_BSS_AVG_ACCESS_DELAY = 63,
	   WLAN_EID_ANTENNA_INFO = 64,
	   WLAN_EID_RSNI = 65,
	   WLAN_EID_MEASUREMENT_PILOT_TX_INFO = 66,
	   WLAN_EID_BSS_AVAILABLE_CAPACITY = 67,
	   WLAN_EID_BSS_AC_ACCESS_DELAY = 68,
	   WLAN_EID_TIME_ADVERTISEMENT = 69,
	   WLAN_EID_RRM_ENABLED_CAPABILITIES = 70,
	   WLAN_EID_MULTIPLE_BSSID = 71,
	   WLAN_EID_BSS_COEX_2040 = 72,
	   WLAN_EID_BSS_INTOLERANT_CHL_REPORT = 73,
	   WLAN_EID_OVERLAP_BSS_SCAN_PARAM = 74,
	   WLAN_EID_RIC_DESCRIPTOR = 75,
	   WLAN_EID_MMIE = 76,
	   WLAN_EID_ASSOC_COMEBACK_TIME = 77,
	   WLAN_EID_EVENT_REQUEST = 78,
	   WLAN_EID_EVENT_REPORT = 79,
	   WLAN_EID_DIAGNOSTIC_REQUEST = 80,
	   WLAN_EID_DIAGNOSTIC_REPORT = 81,
	   WLAN_EID_LOCATION_PARAMS = 82,
	   WLAN_EID_NON_TX_BSSID_CAP =  83,
	   WLAN_EID_SSID_LIST = 84,
	   WLAN_EID_MULTI_BSSID_IDX = 85,
	   WLAN_EID_FMS_DESCRIPTOR = 86,
	   WLAN_EID_FMS_REQUEST = 87,
	   WLAN_EID_FMS_RESPONSE = 88,
	   WLAN_EID_QOS_TRAFFIC_CAPA = 89,
	   WLAN_EID_BSS_MAX_IDLE_PERIOD = 90,
	   WLAN_EID_TSF_REQUEST = 91,
	   WLAN_EID_TSF_RESPOSNE = 92,
	   WLAN_EID_WNM_SLEEP_MODE = 93,
	   WLAN_EID_TIM_BCAST_REQ = 94,
	   WLAN_EID_TIM_BCAST_RESP = 95,
	   WLAN_EID_COLL_IF_REPORT = 96,
	   WLAN_EID_CHANNEL_USAGE = 97,
	   WLAN_EID_TIME_ZONE = 98,
	   WLAN_EID_DMS_REQUEST = 99,
	   WLAN_EID_DMS_RESPONSE = 100,
	   WLAN_EID_LINK_ID = 101,
	   WLAN_EID_WAKEUP_SCHEDUL = 102,
	   // 103 reserved
	   WLAN_EID_CHAN_SWITCH_TIMING = 104,
	   WLAN_EID_PTI_CONTROL = 105,
	   WLAN_EID_PU_BUFFER_STATUS = 106,
	   WLAN_EID_INTERWORKING = 107,
	   WLAN_EID_ADVERTISEMENT_PROTOCOL = 108,
	   WLAN_EID_EXPEDITED_BW_REQ = 109,
	   WLAN_EID_QOS_MAP_SET = 110,
	   WLAN_EID_ROAMING_CONSORTIUM = 111,
	   WLAN_EID_EMERGENCY_ALERT = 112,
	   WLAN_EID_MESH_CONFIG = 113,
	   WLAN_EID_MESH_ID = 114,
	   WLAN_EID_LINK_METRIC_REPORT = 115,
	   WLAN_EID_CONGESTION_NOTIFICATION = 116,
	   WLAN_EID_PEER_MGMT = 117,
	   WLAN_EID_CHAN_SWITCH_PARAM = 118,
	   WLAN_EID_MESH_AWAKE_WINDOW = 119,
	   WLAN_EID_BEACON_TIMING = 120,
	   WLAN_EID_MCCAOP_SETUP_REQ = 121,
	   WLAN_EID_MCCAOP_SETUP_RESP = 122,
	   WLAN_EID_MCCAOP_ADVERT = 123,
	   WLAN_EID_MCCAOP_TEARDOWN = 124,
	   WLAN_EID_GANN = 125,
	   WLAN_EID_RANN = 126,
	   WLAN_EID_EXT_CAPABILITY = 127,
	   // 128, 129 reserved for Agere
	   WLAN_EID_PREQ = 130,
	   WLAN_EID_PREP = 131,
	   WLAN_EID_PERR = 132,
	   // 133-136 reserved for Cisco
	   WLAN_EID_PXU = 137,
	   WLAN_EID_PXUC = 138,
	   WLAN_EID_AUTH_MESH_PEER_EXCH = 139,
	   WLAN_EID_MIC = 140,
	   WLAN_EID_DESTINATION_URI = 141,
	   WLAN_EID_UAPSD_COEX = 142,
	   WLAN_EID_WAKEUP_SCHEDULE = 143,
	   WLAN_EID_EXT_SCHEDULE = 144,
	   WLAN_EID_STA_AVAILABILITY = 145,
	   WLAN_EID_DMG_TSPEC = 146,
	   WLAN_EID_DMG_AT = 147,
	   WLAN_EID_DMG_CAP = 148,
	   // 149 reserved for Cisco
	   WLAN_EID_CISCO_VENDOR_SPECIFIC = 150,
	   WLAN_EID_DMG_OPERATION = 151,
	   WLAN_EID_DMG_BSS_PARAM_CHANGE = 152,
	   WLAN_EID_DMG_BEAM_REFINEMENT = 153,
	   WLAN_EID_CHANNEL_MEASURE_FEEDBACK = 154,
	   // 155-156 reserved for Cisco
	   WLAN_EID_AWAKE_WINDOW = 157,
	   WLAN_EID_MULTI_BAND = 158,
	   WLAN_EID_ADDBA_EXT = 159,
	   WLAN_EID_NEXT_PCP_LIST = 160,
	   WLAN_EID_PCP_HANDOVER = 161,
	   WLAN_EID_DMG_LINK_MARGIN = 162,
	   WLAN_EID_SWITCHING_STREAM = 163,
	   WLAN_EID_SESSION_TRANSITION = 164,
	   WLAN_EID_DYN_TONE_PAIRING_REPORT = 165,
	   WLAN_EID_CLUSTER_REPORT = 166,
	   WLAN_EID_RELAY_CAP = 167,
	   WLAN_EID_RELAY_XFER_PARAM_SET = 168,
	   WLAN_EID_BEAM_LINK_MAINT = 169,
	   WLAN_EID_MULTIPLE_MAC_ADDR = 170,
	   WLAN_EID_U_PID = 171,
	   WLAN_EID_DMG_LINK_ADAPT_ACK = 172,
	   // 173 reserved for Symbol
	   WLAN_EID_MCCAOP_ADV_OVERVIEW = 174,
	   WLAN_EID_QUIET_PERIOD_REQ = 175,
	   // 176 reserved for Symbol
	   WLAN_EID_QUIET_PERIOD_RESP = 177,
	   // 178-179 reserved for Symbol
	   // 180 reserved for ISO/IEC 20011
	   WLAN_EID_EPAC_POLICY = 182,
	   WLAN_EID_CLISTER_TIME_OFF = 183,
	   WLAN_EID_INTER_AC_PRIO = 184,
	   WLAN_EID_SCS_DESCRIPTOR = 185,
	   WLAN_EID_QLOAD_REPORT = 186,
	   WLAN_EID_HCCA_TXOP_UPDATE_COUNT = 187,
	   WLAN_EID_HL_STREAM_ID = 188,
	   WLAN_EID_GCR_GROUP_ADDR = 189,
	   WLAN_EID_ANTENNA_SECTOR_ID_PATTERN = 190,
	   WLAN_EID_VHT_CAPABILITY = 191,
	   WLAN_EID_VHT_OPERATION = 192,
	   WLAN_EID_EXTENDED_BSS_LOAD = 193,
	   WLAN_EID_WIDE_BW_CHANNEL_SWITCH = 194,
	   WLAN_EID_VHT_TX_POWER_ENVELOPE = 195,
	   WLAN_EID_CHANNEL_SWITCH_WRAPPER = 196,
	   WLAN_EID_AID = 197,
	   WLAN_EID_QUIET_CHANNEL = 198,
	   WLAN_EID_OPMODE_NOTIF = 199,

	   WLAN_EID_VENDOR_SPECIFIC = 221,
	   WLAN_EID_QOS_PARAMETER = 222,
	   WLAN_EID_CAG_NUMBER = 237,
	   WLAN_EID_AP_CSN = 239,
	   WLAN_EID_FILS_INDICATION = 240,
	   WLAN_EID_DILS = 241,
	   WLAN_EID_FRAGMENT = 242,
	   WLAN_EID_EXTENSION = 255
	*/

}

type Wireless80211Frame struct {
	Length           uint16
	TSFT             uint64
	FlagsRadio       layers.RadioTapFlags
	Rate             layers.RadioTapRate
	ChannelFrequency layers.RadioTapChannelFrequency
	ChannelFlags     layers.RadioTapChannelFlags
	FHSS             uint16
	DBMAntennaSignal int8
	DBMAntennaNoise  int8
	LockQuality      uint16
	TxAttenuation    uint16
	DBTxAttenuation  uint16
	DBMTxPower       int8
	Antenna          uint8
	DBAntennaSignal  uint8
	DBAntennaNoise   uint8
	RxFlags          layers.RadioTapRxFlags
	TxFlags          layers.RadioTapTxFlags
	RtsRetries       uint8
	DataRetries      uint8
	MCS              layers.RadioTapMCS
	AMPDUStatus      layers.RadioTapAMPDUStatus
	VHT              layers.RadioTapVHT
	Type             string
	Flags80211       layers.Dot11Flags
	Proto            uint8
	DurationID       uint16
	Address1         string
	Address2         string
	Address3         string
	Address4         string
	SequenceNumber   uint16
	FragmentNumber   uint16
	Checksum         uint32
	Elements         map[string][]byte
	Interface        string
}

func (frame *Wireless80211Frame) ParseElements(packet gopacket.Packet, ether *layers.Dot11) {
	if probegreq, ok := packet.Layer(layers.LayerTypeDot11MgmtProbeReq).(*layers.Dot11MgmtProbeReq); ok {
		frame.Elements = ParseFrameElements(probegreq.LayerContents())
	} else if beaconframe, ok := packet.Layer(layers.LayerTypeDot11MgmtBeacon).(*layers.Dot11MgmtBeacon); ok {
		frame.Elements = ParseFrameElements(beaconframe.LayerContents()[12:])
	} else if _, ok := packet.Layer(layers.LayerTypeDot11Data).(*layers.Dot11Data); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataCFAck).(*layers.Dot11DataCFAck); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataCFAckNoData).(*layers.Dot11DataCFAckNoData); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataCFAckPoll).(*layers.Dot11DataCFAckPoll); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataCFAckPollNoData).(*layers.Dot11DataCFAckPollNoData); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataCFPoll).(*layers.Dot11DataCFPoll); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataCFPollNoData).(*layers.Dot11DataCFPollNoData); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataNull).(*layers.Dot11DataNull); ok {
		//} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOS).(*layers.Dot11DataQOS); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOSCFAckPollNoData).(*layers.Dot11DataQOSCFAckPollNoData); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOSCFPollNoData).(*layers.Dot11DataQOSCFPollNoData); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOSData).(*layers.Dot11DataQOSData); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOSDataCFAck).(*layers.Dot11DataQOSDataCFAck); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOSDataCFAckPoll).(*layers.Dot11DataQOSDataCFAckPoll); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOSDataCFPoll).(*layers.Dot11DataQOSDataCFPoll); ok {
	} else if _, ok := packet.Layer(layers.LayerTypeDot11DataQOSNull).(*layers.Dot11DataQOSNull); ok {
	}
}

func ParseFrameElements(stream []byte) (elements map[string][]byte) {
	elements = map[string][]byte{}
	for len(stream) > 0 {
		field_id, remainder := stream[0], stream[1:]
		stream = remainder

		field, ok := ELEMENT_IDS[field_id]
		if !ok {
			log.WithFields(log.Fields{
				"at": "models.ParseFrameElements",
				"id": field_id,
			}).Warn("unknown element id")
			return
		}

		field_len, remainder := stream[0], stream[1:]
		stream = remainder
		if field_len == 0 {
			continue
		}

		field_data, remainder := stream[:field_len], stream[field_len:]
		stream = remainder

		elements[field] = field_data
	}
	return
}
