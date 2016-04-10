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

var version = flag.Bool("version", false, "version")

func main() {
	flag.Parse()

	if *version {
		fmt.Println("Wave 0.0.0")
		os.Exit(0)
	}

	// if reseed db

	database.SetupElasticsearch()
	database.DB()

	if helpers.Production() {
		log.SetFormatter(&log.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
		controllers.NewRouter().Run(":80")
	} else if helpers.Development() {
		controllers.NewRouter().Run(":8080")
	}
}
