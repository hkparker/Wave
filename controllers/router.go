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
	router.POST("/login", newSession)
	//router.POST("/logout", deleteSession)

	// User routes
	router.POST("/users/create", createUser)
	router.POST("/users/name", updateUserName)
	//router.GET("/users/password/:id", passwordReset)
	router.POST("/users/password", updateUserPassword)
	router.POST("/users/destroy", deleteUser)

	// Signature routes
	// Incident routes
	// Device routes

	// Collector routes
	router.GET("/admin/collectors", getCollectors)
	router.POST("/admin/collectors/create", createCollector)
	router.POST("/admin/collectors/delete", deleteCollector)

	// Certificate Management
	router.GET("/admin/tls", getTLS)
	router.POST("/admin/tls", setTLS)

	return router
}

func NewCollector() *gin.Engine {
	collector := gin.Default()
	collector.GET("/frames", pollCollector)
	return collector
}
