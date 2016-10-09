package models

import (
	"crypto/tls"
	"crypto/x509"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/database"
	"github.com/jinzhu/gorm"
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
		Certificates: APITLSConfig().Certificates,
		ClientCAs:    cert_pool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	config.BuildNameToCertificate()
	return config
}

func newCollector() {
	// Generate new client certificate and save
}
