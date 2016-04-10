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
	if helpers.Production() {
		log.SetFormatter(&log.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	}
	database.SetupElasticsearch()
	database.DB()

	controllers.NewRouter(database.DB()).Run(":8080")
}
