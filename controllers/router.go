package controllers

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/middleware"
)

func NewAPI() *gin.Engine {
	router := gin.Default()
	if helpers.Production {
		router.Use(middleware.EmbeddedAssets())
	} else {
		router.Use(static.Serve("/", static.LocalFile("static", false)))
		router.GET("/", middleware.RenderWebpack)
	}
	router.Use(middleware.Authentication())

	// Session routes
	router.POST("/sessions/create", newSession)
	router.POST("/sessions/delete", deleteSession)

	// User routes
	router.GET("/users", getUsers)
	router.POST("/users/create", createUser)
	router.POST("/users/name", updateUserName)
	router.POST("/users/password", updateUserPassword)
	router.POST("/users/assign-password", assignUserPassword)
	router.POST("/users/delete", deleteUser)

	// Frontend Events

	// Visualizer stream

	// Signature routes
	// Incident routes
	// Device routes
	// Version route

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
	collector := gin.Default()
	collector.GET("/frames", pollCollector)
	return collector
}
