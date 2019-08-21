package controllers

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	//"github.com/hkparker/Wave/engines/visualizer"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/middleware"
	cors "github.com/rs/cors/wrapper/gin"
)

func NewAPI() *gin.Engine {
	router := gin.Default()

	if helpers.Development { // create a proper CORS for production
		c := cors.New(cors.Options{
		    AllowedOrigins: []string{"http://localhost:8080"},
		    AllowCredentials: true,
		    //Debug: true,
		})
		router.Use(c)
	}

	if helpers.Production {
		router.Use(middleware.EmbeddedAssets())
	} else {
		router.Use(static.Serve("/", static.LocalFile("static", false)))
		router.GET("/", middleware.RenderWebpack)
	}
	router.Use(middleware.Authentication())

	// Session routes
	router.POST("/sessions/create", newSession)
	router.POST("/sessions/destroy", deleteSession)

	// User routes
	router.GET("/users", getUsers)
	router.POST("/users/create", createUser)
	router.POST("/users/name", updateUserName)
	router.POST("/users/password", updateUserPassword)
	router.POST("/users/assign-password", assignUserPassword)
	router.POST("/users/delete", deleteUser)

	// Event streams
	router.GET("/streams/visualizer", streamVisualization)
	//router.GET("/streams/ids", streamAlerts)

	// Signature routes
	// Incident routes
	// Device routes
	// Version route

	router.GET("/version", func (c *gin.Context) { c.JSON(200, gin.H{"version": 0}) })

	// Collector routes
	router.GET("/collectors", getCollectors)
	router.POST("/collectors/create", createCollector)
	router.POST("/collectors/delete", deleteCollector)

	// Certificate Management
	router.GET("/tls", getTLS)
	router.POST("/tls", setTLS)

	return router
}

func NewCollector() *gin.Engine {
	//visualizer.Load()

	collector := gin.Default()
	collector.GET("/frames", pollCollector)
	return collector
}
