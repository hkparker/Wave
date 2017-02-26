package helpers

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"strconv"
)

var WaveAddress string
var WaveHost string

func SetHostname() {
	address := os.Getenv("WAVE_ADDRESS")
	if address == "" {
		log.Fatal("WAVE_ADDRESS envar must be provided")
	}
	port_str := os.Getenv("WAVE_PORT")
	var port int
	if address == "" {
		log.Fatal("WAVE_PORT envar must be provided")
	} else {
		var err error
		port, err = strconv.Atoi(port_str)
		if err != nil {
			log.WithFields(log.Fields{
				"at":    "helpers.SetHostname",
				"value": port_str,
			}).Fatal("unable to assert WAVE_PORT as int")
		}
	}
	tls := false
	if os.Getenv("WAVE_TLS") == "true" {
		tls = true
	}

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
