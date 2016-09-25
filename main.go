package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/controllers"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
	"os"
)

func main() {
	var version bool
	var initdb bool
	var port int
	var address string
	var db_username string
	var db_password string
	var db_name string
	var db_ssl string
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&initdb, "initdb", false, "reset the Wave database")
	flag.IntVar(&port, "port", 80, "port to listen on")
	flag.StringVar(&address, "address", "0.0.0.0", "ip address to bind to")
	flag.StringVar(&db_username, "db_username", "", "username for Wave database")
	flag.StringVar(&db_password, "db_password", "", "password for Wave database")
	flag.StringVar(&db_name, "db_name", "wave_development", "database name to use")
	flag.StringVar(&db_ssl, "db_ssl", "disabled", "database connection over ssl")
	flag.Parse()

	if version {
		fmt.Println("Wave 0.0.0")
		os.Exit(0)
	}

	database.Connect(
		db_username,
		db_password,
		db_name,
		db_ssl,
	)
	//cache.Connect()

	if helpers.Production() {
		log.SetFormatter(&log.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	}
	controllers.NewRouter().Run(fmt.Sprintf(
		"%s:%d",
		address,
		port,
	))
}
