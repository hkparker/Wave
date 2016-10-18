package models

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
