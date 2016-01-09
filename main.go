package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/controllers"
	"net/http"
)

func getting(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.tmpl", gin.H{})
}

func renderLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func render2FA(c *gin.Context) {
	c.HTML(http.StatusOK, "2fa.tmpl", gin.H{})
}

func renderReset(c *gin.Context) {
	c.HTML(http.StatusOK, "reset.tmpl", gin.H{})
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/vendor", static.LocalFile("vendor", false)))
	router.Use(static.Serve("/", static.LocalFile("static", false)))
	router.LoadHTMLGlob("views/*")
	router.GET("/", getting)
	router.GET("/login", renderLogin)
	router.GET("/2fa", render2FA)
	router.GET("/reset", renderReset)
	router.POST("/frame", controllers.FrameFromCollector)
	return router
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	NewRouter().Run(":8080")
}
