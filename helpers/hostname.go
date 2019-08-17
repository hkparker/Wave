package helpers

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var WaveAddress string
var WaveBind string
var WaveHostname string
var WavePort int
var CollectorPort int
var TLS bool

func init() {
	if TestingCmd() {
		WaveAddress = "http://127.0.0.1"
		WaveBind = "127.0.0.1"
		WaveHostname = "localhost"
		WavePort = 8080
		CollectorPort = 8888
		TLS = false
	}
}

func setHostname() {
	WaveHostname = os.Getenv("WAVE_HOSTNAME")
	if WaveHostname == "" {
		log.Fatal("WAVE_HOSTNAME envar must be provided")
	}
	WaveBind = os.Getenv("WAVE_BIND")
	if WaveBind == "" {
		log.Fatal("WAVE_BIND envar must be provided")
	}

	port_str := os.Getenv("WAVE_PORT")
	collector_port_str := os.Getenv("WAVE_COLLECTOR_PORT")
	if port_str == "" || collector_port_str == "" {
		log.Fatal("WAVE_PORT and WAVE_COLLECTOR_PORT envars must be provided")
	} else {
		var err error
		WavePort, err = strconv.Atoi(port_str)
		if err != nil {
			log.WithFields(log.Fields{
				"at":    "helpers.SetHostname",
				"value": port_str,
			}).Fatal("unable to assert WAVE_PORT as int")
		}
		CollectorPort, err = strconv.Atoi(collector_port_str)
		if err != nil {
			log.WithFields(log.Fields{
				"at":    "helpers.SetHostname",
				"value": port_str,
			}).Fatal("unable to assert WAVE_COLLECTOR_PORT as int")
		}
	}

	TLS = false
	if os.Getenv("WAVE_TLS") == "true" {
		TLS = true
	}

	if TLS {
		WaveAddress = "https://"
	} else {
		WaveAddress = "http://"
	}
	WaveAddress = WaveAddress + WaveHostname
	if !((!TLS && (WavePort == 80)) || (TLS && (WavePort == 443))) {
		WaveAddress = WaveAddress + ":" + strconv.Itoa(WavePort)
	}
}
