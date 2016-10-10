package helpers

import (
	"strconv"
)

var WaveAddress string
var WaveHost string

func SetHostname(tls bool, address string, port int) {
	WaveHost = address
	if tls {
		WaveAddress = "https://"
	} else {
		WaveAddress = "http://"
	}
	WaveAddress = WaveAddress + address
	if !((!tls && (port == 80)) || (tls && (port == 443))) {
		WaveAddress = WaveAddress + ":" + strconv.Itoa(port)
	}
}
