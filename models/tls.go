package models

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/jinzhu/gorm"
	"math/big"
	"net"
	"time"
)

type TLS struct {
	gorm.Model
	CaCert     string
	PrivateKey string
}

func APITLSConfig() (config *tls.Config) {
	config = &tls.Config{
		Certificates: []tls.Certificate{APITLSCertificate()},
	}
	return
}

func APITLSCertificate() (pair tls.Certificate) {
	ca_cert, private_key := APITLSData()
	var err error
	pair, err = tls.X509KeyPair(
		ca_cert,
		private_key,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.APITLSCertificate",
			"error": err.Error(),
		}).Fatal("error reading TLS data from database")
	}
	return
}

func APITLSData() ([]byte, []byte) {
	createTLSIfMissing()
	var model TLS
	err := Orm.First(&model).Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.APITLSData",
			"error": err.Error(),
		}).Fatal("error loading first TLS record for data")
	}
	return []byte(model.CaCert), []byte(model.PrivateKey)
}

func SetTLS(request map[string]string) (err error) {
	createTLSIfMissing()
	var collectors []Collector
	var collector_count int
	err = Orm.Find(&collectors).Count(&collector_count).Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.SetTLS",
			"error": err.Error(),
		}).Error("unable to load collector count for setting tls")
		return
	}
	if collector_count != 0 {
		err_str := "cannot set TLS data until all collectors are deleted"
		log.WithFields(log.Fields{
			"at": "models.SetTLS",
		}).Error(err_str)
		err = errors.New(err_str)
		return
	}
	var config TLS
	err = Orm.First(&config).Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.SetTLS",
			"error": err.Error(),
		}).Error("error loading TLS config to set")
		return
	}

	if _, ok := request["ca_cert"]; !ok {
		err_str := "missing ca_cert key"
		log.WithFields(log.Fields{
			"at": "models.SetTLS",
		}).Error(err_str)
		err = errors.New(err_str)
		return
	}

	if _, ok := request["private_key"]; !ok {
		err_str := "missing private_key key"
		log.WithFields(log.Fields{
			"at": "models.SetTLS",
		}).Error(err_str)
		err = errors.New(err_str)
		return
	}

	config.CaCert = request["ca_cert"]
	config.PrivateKey = request["private_key"]
	err = Orm.Save(&config).Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.SetTLS",
			"error": err.Error(),
		}).Error("error saving new TLS data")
	}
	return
}

//
// Used to generate a new self-signed certificate if there is no
// certificate in the database when loading certificate data.
//
func createTLSIfMissing() (err error) {
	var count int
	var tls []TLS
	err = Orm.Find(&tls).Count(&count).Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.createTLSIfMissing",
			"error": err.Error(),
		}).Fatal("error loading tls count from database")
	}
	if count == 0 {
		cert_data, key_data := selfSignedCert()
		new_config := TLS{
			CaCert:     string(cert_data),
			PrivateKey: string(key_data),
		}
		err = Orm.Save(&new_config).Error
		if err != nil {
			log.WithFields(log.Fields{
				"at":    "models.createTLSIfMissing",
				"error": err.Error(),
			}).Fatal("error saving new self-signed certificate")
		}
	}
	return
}

func randomSerial() *big.Int {
	serial_number_limit := new(big.Int).Lsh(big.NewInt(1), 128)
	serial_number, err := rand.Int(rand.Reader, serial_number_limit)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.randomSerial",
			"error": err.Error(),
		}).Fatal("failed to generate serial number for self signed certificate")
	}
	return serial_number
}

// Generate a new self-signed certificate to be used if the --tls
// flag is set but no TLS certificate and key are stored in the
func selfSignedCert() (cert_data []byte, key_data []byte) {
	// Self signed certificate for provided hostname
	ca := &x509.Certificate{
		SerialNumber: randomSerial(),
		Subject: pkix.Name{
			Organization: []string{"Wave"},
			CommonName:   helpers.WaveHostname,
		},
		IPAddresses:           []net.IP{net.ParseIP(helpers.WaveBind)},
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
		log.WithFields(log.Fields{
			"at":    "models.selfSignedCert",
			"error": err.Error(),
		}).Fatal("failed to generate private key for self signed certificate")
	}
	pub := &priv.PublicKey

	// Create Certificate
	cert_der, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.selfSignedCert",
			"error": err.Error(),
		}).Fatal("failed to create self signed certificate")
	}

	// Create PEM encoding of certificate
	var cert_buffer bytes.Buffer
	err = pem.Encode(&cert_buffer, &pem.Block{Type: "CERTIFICATE", Bytes: cert_der})
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.selfSignedCert",
			"error": err.Error(),
		}).Fatal("could not PEM encode certificate data")
	}
	cert_data = cert_buffer.Bytes()

	// Create PEM encoding of key
	var key_buffer bytes.Buffer
	err = pem.Encode(&key_buffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "models.selfSignedCert",
			"error": err.Error(),
		}).Fatal("could not PEM encode key data")
	}
	key_data = key_buffer.Bytes()

	return
}
