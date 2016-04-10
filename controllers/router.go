package controllers

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	if helpers.Production() {
		router.Use(middleware.EmbeddedAssets())
	} else {
		router.Use(static.Serve("/", static.LocalFile("static", false)))
		router.GET("/", middleware.RenderWebpack)
	}
	router.Use(middleware.Authentication())

	// Authentication routes
	router.POST("/login", Login)
	router.GET("/reset/:id", PasswordReset)

	// Users routes
	router.POST("/users/create", CreateUser)

	// Collector
	router.GET("/frames", PollCollector)

	return router
}
