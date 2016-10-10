package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
)

func getTLS(c *gin.Context) {
	cert_data, key_data := models.APITLSData()
	c.JSON(
		200,
		gin.H{
			"certificate": string(cert_data),
			"private_key": string(key_data),
		},
	)
}

func setTLS(c *gin.Context) {
	tls_info, err := requestJSON(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = models.SetTLS(tls_info)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
