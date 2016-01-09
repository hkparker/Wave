package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FrameFromCollector(c *gin.Context) {
	c.HTML(http.StatusOK, "reset.tmpl", gin.H{})
}
