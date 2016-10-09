package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
)

func getTLS(c *gin.Context) {
	//config, err := models.APITLSConfig()
	//if err != nil {
	//}
	c.JSON(200, "") //config)
}

func setTLS(c *gin.Context) {
	tls_info, err := requestJSON(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = models.SetTLS(tls_info)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
