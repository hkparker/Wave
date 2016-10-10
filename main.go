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
	"net"
	"net/http"
	"os"
)

func main() {
	var version bool
	var initdb bool
	var port int
	var collector_port int
	var address string
	var api_tls bool
	var db_username string
	var db_password string
	var db_name string
	var db_ssl string
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&initdb, "initdb", false, "reset the Wave database")
	flag.IntVar(&port, "port", 80, "port to listen on")
	flag.IntVar(&collector_port, "collector-port", 444, "port to listen for collector websockets on")
	flag.StringVar(&address, "address", "0.0.0.0", "ip address to bind to")
	flag.BoolVar(&api_tls, "tls", false, "serve over TLS socket")
	flag.StringVar(&db_username, "db_username", "", "username for Wave database")
	flag.StringVar(&db_password, "db_password", "", "password for Wave database")
	flag.StringVar(&db_name, "db_name", "wave_development", "database name to use")
	flag.StringVar(&db_ssl, "db_ssl", "disable", "database connection over ssl")
	flag.Parse()

	if version {
		fmt.Println("Wave 0.0.0")
		os.Exit(0)
	}

	helpers.SetHostname(api_tls, address, port)
	helpers.SetEnvironment()
	database.Connect(
		db_username,
		db_password,
		db_name,
		db_ssl,
	)
	models.CreateTables()

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
			}).Fatal("unable to create listener for connector")
		}
		tls_listener := tls.NewListener(tcp_listener, config)
		server.Serve(tls_listener)
	}

	// Start Collector server
	go run_tls(
		controllers.NewCollector(),
		fmt.Sprintf(
			"%s:%d",
			address,
			collector_port,
		),
		models.CollectorTLSConfig(),
	)

	// Start Wave API
	if api_tls {
		run_tls(
			controllers.NewAPI(),
			fmt.Sprintf(
				"%s:%d",
				address,
				port,
			),
			models.APITLSConfig(),
		)
	} else {
		controllers.NewAPI().Run(fmt.Sprintf(
			"%s:%d",
			address,
			port,
		))
	}
}
