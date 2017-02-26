package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/controllers"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	_ "github.com/joho/godotenv/autoload"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var version bool
	var initdb bool
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&initdb, "initdb", false, "reset the Wave database")
	flag.Parse()

	if version {
		fmt.Println("Wave 0.0.0")
		os.Exit(0)
	}

	helpers.SetHostname()
	helpers.SetEnvironment()
	database.Connect(
		os.Getenv("WAVE_DB_USERNAME"),
		os.Getenv("WAVE_DB_PASSWORD"),
		os.Getenv("WAVE_DB_NAME"),
		os.Getenv("WAVE_DB_TLS"),
	)
	models.CreateTables()
	models.CreateAdmin()

	// Create small wrapper for starting http handlers with TLS configs
	run_tls := func(handler http.Handler, address string, config *tls.Config) {
		server := &http.Server{
			Handler: handler,
		}
		tcp_listener, err := net.Listen("tcp", address)
		if err != nil {
			log.WithFields(log.Fields{
				"address": address,
				"error":   err.Error(),
			}).Fatal("unable to create tls listener")
		}
		tls_listener := tls.NewListener(tcp_listener, config)
		server.Serve(tls_listener)
	}

	// Parse port envars
	collector_port_str := os.Getenv("WAVE_COLLECTOR_PORT")
	wave_port_str := os.Getenv("WAVE_PORT")
	var collector_port int
	var wave_port int
	if collector_port_str == "" || wave_port_str == "" {
		log.Fatal("WAVE_PORT and WAVE_COLLECTOR_PORT envars must be provided")
	} else {
		var err error
		wave_port, err = strconv.Atoi(wave_port_str)
		if err != nil {
			log.WithFields(log.Fields{
				"at":    "main",
				"value": wave_port_str,
			}).Fatal("unable to assert WAVE_PORT as int")
		}
		collector_port, err = strconv.Atoi(collector_port_str)
		if err != nil {
			log.WithFields(log.Fields{
				"at":    "main",
				"value": collector_port_str,
			}).Fatal("unable to assert WAVE_COLLECTOR_PORT as int")
		}
	}

	// Start Collector server
	go run_tls(
		controllers.NewCollector(),
		fmt.Sprintf(
			"%s:%d",
			os.Getenv("WAVE_ADDRESS"),
			collector_port,
		),
		models.CollectorTLSConfig(),
	)

	// Start Wave API
	if os.Getenv("WAVE_TLS") == "true" {
		run_tls(
			controllers.NewAPI(),
			fmt.Sprintf(
				"%s:%d",
				os.Getenv("WAVE_ADDRESS"),
				os.Getenv("WAVE_PORT"),
			),
			models.APITLSConfig(),
		)
	} else {
		controllers.NewAPI().Run(fmt.Sprintf(
			"%s:%d",
			os.Getenv("WAVE_ADDRESS"),
			wave_port,
		))
	}
}
