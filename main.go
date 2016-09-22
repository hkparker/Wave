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
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&initdb, "initdb", false, "reset the Wave database")
	flag.Parse()

	if version {
		fmt.Println("Wave 0.0.0")
		os.Exit(0)
	}

	database.Connect()
	//database.SetupElasticsearch()

	if initdb {
		database.Init()
	}

	if helpers.Production() {
		log.SetFormatter(&log.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
		controllers.NewRouter().Run(":80")
	} else if helpers.Development() {
		controllers.NewRouter().Run(":8080")
	}
}
