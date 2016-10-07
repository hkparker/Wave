package models

import (
	log "github.com/Sirupsen/logrus"
	//"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/database"
	//"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
	//"time"
	"crypto/tls"
	"crypto/x509"
)

type Collector struct {
	gorm.Model
	Name   string
	CaCert []byte
}

func CollectorTLSConfig() *tls.Config {
	// Load all collectors
	var collectors []Collector
	err := database.Orm.Find(&collectors).Error
	if err != nil {
		log.Fatal(err)
	}

	// Create certificate pool for collectors
	cert_pool := x509.NewCertPool()
	for _, collector := range collectors {
		client_cert := collector.CaCert
		cert_pool.AppendCertsFromPEM(client_cert)
	}

	// Create client validating TLS config and return
	config := &tls.Config{
		ClientCAs:  cert_pool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	config.BuildNameToCertificate()
	return config
}
