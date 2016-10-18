package models

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TextCollectorTLSConfigCreatesCorrectConfig(t *testing.T) {
	assert := assert.New(t)

	collector1, err := CreateCollector("thing1")
	assert.Nil(err)
	collector2, err := CreateCollector("thing2")
	assert.Nil(err)

	config := CollectorTLSConfig()
	if assert.Equal(2, len(config.Certificates)) {
		assert.NotEqual(config.Certificates[0], config.Certificates[1])
	}

	block1, _ := pem.Decode(collector1.CaCert)
	client1, err := x509.ParseCertificate(block1.Bytes)
	assert.Nil(err)

	block2, _ := pem.Decode(collector2.CaCert)
	client2, err := x509.ParseCertificate(block2.Bytes)
	assert.Nil(err)

	for _, cert := range config.Certificates {
		assert.Equal(
			true,
			bytes.Equal(
				cert.Certificate[0],
				client1.Raw,
			) || bytes.Equal(
				cert.Certificate[0],
				client2.Raw,
			),
		)
	}
}

func TestNewCollectorKeysAreSignedByAPIKey(t *testing.T) {
	assert := assert.New(t)

	collector_cert_data, _, err := newCollectorKeys()
	assert.Nil(err)
	block, _ := pem.Decode(collector_cert_data)
	ca, err := x509.ParseCertificate(APITLSCertificate().Certificate[0])
	assert.Nil(err)
	client, err := x509.ParseCertificate(block.Bytes)
	assert.Nil(err)

	err = client.CheckSignatureFrom(ca)
	assert.Nil(err)
}
