package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/controllers"
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/middleware"
	"os"
)

var version = flag.Bool("version", false, "version")

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.EmbeddedAssets())
	router.Use(middleware.Authentication(database.DB()))

	// Authentication routes
	router.POST("/login", controllers.Login)
	router.POST("/2fa", controllers.SubmitTwoFactor)
	router.GET("/reset/:id", controllers.PasswordReset)

	// Collector
	router.GET("/frames", controllers.PollCollector)

	return router
}

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
	NewRouter().Run(":8080")
}
