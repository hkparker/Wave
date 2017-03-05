package models

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"time"
)

type Collector struct {
	gorm.Model
	Name       string `sql:"not null;unique"`
	CaCert     string
	PrivateKey string
}

func CollectorTLSConfig() *tls.Config {
	// Validate the client was signed by the Wave API certificate
	cert_pool := x509.NewCertPool()
	cert, _ := APITLSData()
	cert_pool.AppendCertsFromPEM(cert)

	// Create client validating TLS config and return
	config := &tls.Config{
		Certificates: APITLSConfig().Certificates,
		ClientCAs:    cert_pool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	config.BuildNameToCertificate()
	return config
}

func Collectors() (collectors []Collector, err error) {
	err = Orm.Find(&collectors).Error
	return
}

func CreateCollector(name string) (collector Collector, err error) {
	cert_data, key_data, err := newCollectorKeys()
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.CreateCollector",
			"error": err.Error(),
		}).Error("failed to create collector")
		return
	}
	collector = Collector{
		Name:       name,
		CaCert:     string(cert_data),
		PrivateKey: string(key_data),
	}
	err = Orm.Save(&collector).Error
	return
}

func newCollectorKeys() (cert_data []byte, key_data []byte, err error) {
	api_cert := APITLSCertificate()
	ca, err := x509.ParseCertificate(api_cert.Certificate[0])
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.newColletorKeys",
			"error": err.Error(),
		}).Error("error parsing API TLS certificate for new collector")
		return
	}
	ca_key := api_cert.PrivateKey

	collector_cert := &x509.Certificate{
		SerialNumber: randomSerial(),
		Subject: pkix.Name{
			Organization:       []string{"Wave"},
			OrganizationalUnit: []string{"Wave"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(6, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	collector_priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.newColletorKeys",
			"error": err.Error(),
		}).Error("error generating private key for new collector")
		return
	}
	collector_pub := &collector_priv.PublicKey
	collector_cert_data, err := x509.CreateCertificate(rand.Reader, collector_cert, ca, collector_pub, ca_key)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.newColletorKeys",
			"error": err.Error(),
		}).Error("error creating collector certificate")
		return
	}

	// Create PEM encoding of certificate
	var cert_buffer bytes.Buffer
	err = pem.Encode(&cert_buffer, &pem.Block{Type: "CERTIFICATE", Bytes: collector_cert_data})
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.newColletorKeys",
			"error": err.Error(),
		}).Error("could not PEM encode collector certificate data")
		return
	}
	cert_data = cert_buffer.Bytes()

	// Create PEM encoding of key
	var key_buffer bytes.Buffer
	err = pem.Encode(&key_buffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(collector_priv)})
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.newColletorKeys",
			"error": err.Error(),
		}).Error("could not PEM encode collector key data")
		return
	}
	key_data = key_buffer.Bytes()

	return
}

func (collector *Collector) Delete() error {
	return Orm.Delete(&collector).Error
}

func CollectorByName(name string) (collector Collector, err error) {
	err = Orm.First(&collector, "Name = ?", name).Error
	return
}
