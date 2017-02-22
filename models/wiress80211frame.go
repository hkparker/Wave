package models

import (
	"github.com/google/gopacket/layers"
)

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
