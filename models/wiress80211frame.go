package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var ELEMENT_IDS = map[byte]string{
	0:   "SSID",
	1:   "SUPPORTED_RATES",
	50:  "EXTENDED_SUPPORTED_RATES",
	3:   "DS_PARAMETER_SET",
	45:  "HT_CAPABILITIES",
	127: "EXTENDED_CAPABILITIES",
	221: "VENDOR_SPECIFIC",
	//107: "??",
	//191: "??",
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
	// since ether.Type is known, try to lookup the right one
	//fmt.Println(ether.Type, uint8(ether.Type))
	//fmt.Println(packet.Layers())
	switch ether.Type {
	case 0: //layers.LayerTypeDot11MgmtBeacon:
		//fmt.Println("its a mgt")
		//if beacon, ok := packet.Layer(layers.LayerTypeDot11MgmtBeacon).(*layers.Dot11MgmtBeacon); ok {
		//	fmt.Println("got a beacon", beacon)
		//	elements = ParseBeaconFrame()
		//}
	}
	if probegreq, ok := packet.Layer(layers.LayerTypeDot11MgmtProbeReq).(*layers.Dot11MgmtProbeReq); ok {
		//fmt.Println(ether.NextLayerType())
		frame.Elements = ParseFrameElements(probegreq.LayerContents())
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
