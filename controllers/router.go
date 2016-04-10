package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/middleware"
	"github.com/jinzhu/gorm"
)

func NewRouter(database *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.EmbeddedAssets())
	router.Use(middleware.Authentication(database))

	// Authentication routes
	router.POST("/login", Login)
	router.GET("/reset/:id", PasswordReset)

	// Users routes
	router.POST("/users/create", CreateUser)

	// Collector
	router.GET("/frames", PollCollector)

	return router
}
