package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.tmpl", gin.H{})
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/vendor", static.LocalFile("vendor", false)))
	router.LoadHTMLGlob("views/*")
	router.GET("/", getting)
	return router
}

func main() {
	//	gin.SetMode(gin.ReleaseMode)
	NewRouter().Run(":8080")
}
