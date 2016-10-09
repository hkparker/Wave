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
		cert_data, key_data := defaultCA()
		new_config := TLS{
			CaCert:     cert_data,
			PrivateKey: key_data,
		}
		database.Orm.Save(&new_config)
	}
	return
}

func defaultCA() (cert_data []byte, key_data []byte) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1500),
		Subject: pkix.Name{
			Country:            []string{"Earth"},
			Organization:       []string{"Wave"},
			OrganizationalUnit: []string{"Wave"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(6, 0, 0),
		BasicConstraintsValid: true,
		IsCA:        true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatal("")
	}
	pub := &priv.PublicKey
	cert_der, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Println("create ca failed", err)
		return
	}
	var cert_buffer bytes.Buffer
	err = pem.Encode(&cert_buffer, &pem.Block{Type: "CERTIFICATE", Bytes: cert_der})
	if err != nil {
		log.Fatal("")
	}
	cert_data = cert_buffer.Bytes()

	var key_buffer bytes.Buffer
	err = pem.Encode(&key_buffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	if err != nil {
		log.Fatal("")
	}
	key_data = key_buffer.Bytes()
	return
}
