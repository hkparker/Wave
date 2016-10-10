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
	"github.com/hkparker/Wave/database"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	"math/big"
	"time"
)

type TLS struct {
	gorm.Model
	CaCert     []byte
	PrivateKey []byte
}

func APITLSConfig() (config *tls.Config) {
	createTLSIfMissing()
	var model TLS
	err := database.Orm.First(&model).Error
	if err != nil {
		log.Fatal("")
	}
	pair, err := tls.X509KeyPair(
		model.CaCert,
		model.PrivateKey,
	)
	if err != nil {
		log.Fatal("")
	}
	config = &tls.Config{
		Certificates: []tls.Certificate{pair},
	}
	return
}

func SetTLS(request map[string]string) (err error) {
	createTLSIfMissing()
	// ensure no collectors
	var config TLS
	err = database.Orm.First(&config).Error
	if err != nil {

	}
	config.CaCert = []byte(request["ca_cert"])
	config.PrivateKey = []byte(request["private_key"])
	err = database.Orm.Save(&config).Error
	if err != nil {

	}
	return
}

func createTLSIfMissing() (err error) {
	var count int
	var tls []TLS
	err = database.Orm.Find(&tls).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		cert_data, key_data := selfSignedCert()
		new_config := TLS{
			CaCert:     cert_data,
			PrivateKey: key_data,
		}
		database.Orm.Save(&new_config)
	}
	return
}

// Generate a new self-signed certificate to be used if the --tls
// flag is set but no TLS certificate and key are stored in the database.
func selfSignedCert() (cert_data []byte, key_data []byte) {
	// Random serial number
	serial_number_limit := new(big.Int).Lsh(big.NewInt(1), 128)
	serial_number, err := rand.Int(rand.Reader, serial_number_limit)
	if err != nil {
		log.Fatalf("failed to generate serial number for self signed certificate: %s", err)
	}

	// Self signed certificate for provided hostname
	ca := &x509.Certificate{
		SerialNumber: serial_number,
		Subject: pkix.Name{
			Organization:       []string{"Wave"},
			OrganizationalUnit: []string{"Wave"},
			CommonName:         helpers.WaveHost,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(6, 0, 0),
		BasicConstraintsValid: true,
		IsCA:        true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// Generate key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate private key for self signed certificate: %s", err)
	}
	pub := &priv.PublicKey

	// Create Certificate
	cert_der, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Fatalf("failed to create self signed certificate: %s", err)
	}

	// Create PEM encoding of certificate
	var cert_buffer bytes.Buffer
	err = pem.Encode(&cert_buffer, &pem.Block{Type: "CERTIFICATE", Bytes: cert_der})
	if err != nil {
		log.Fatal("")
	}
	cert_data = cert_buffer.Bytes()

	// Create PEM encoding of key
	var key_buffer bytes.Buffer
	err = pem.Encode(&key_buffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	if err != nil {
		log.Fatal("")
	}
	key_data = key_buffer.Bytes()

	return
}
