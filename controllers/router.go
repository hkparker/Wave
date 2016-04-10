package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.EmbeddedAssets()) // if in development, serve static from /static as /
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
