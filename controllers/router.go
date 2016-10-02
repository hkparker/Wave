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

	// Session routes
	router.POST("/login", newSession)

	// User routes
	router.POST("/users/create", createUser)
	router.POST("/users/name", updateUserName)
	//router.GET("/users/password/:id", passwordReset)
	router.POST("/users/password", updateUserPassword)
	router.POST("/users/destroy", deleteUser)

	// Collector
	router.GET("/frames", PollCollector)

	return router
}
