package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/controllers"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/middleware"
	"github.com/hkparker/Wave/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

type WaveConfig struct {
	DB gorm.DB
}

var wave = WaveConfig{}
var help = flag.Bool("help", false, "display help message")

func initDB() {
	db, err := gorm.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to postgres")
	}
	models.SetDB(db)
	wave.DB = db
}

func renderWebpack(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	c.String(200,
		`<html>
	<head>
		<meta charset="utf-8">
		<title>Wave</title>
	</head>
	<body>
		<div id="content"></div>
		<script type="text/javascript" src="bundle.js" charset="utf-8"></script>
	</body>
</html>
`,
	)
}

func NewRouter() *gin.Engine {
	log.SetFormatter(&log.JSONFormatter{})
	router := gin.Default()
	router.Use(middleware.Authentication(wave.DB))
	router.Use(static.Serve("/", static.LocalFile("static", false)))
	router.GET("/", renderWebpack)

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
	if *help {
		fmt.Println("Wave 0.0.1")
		os.Exit(0)
	}
	//gin.SetMode(gin.ReleaseMode)
	helpers.SetupElasticsearch()
	initDB()
	NewRouter().Run(":8080")
}
