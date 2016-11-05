package models

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultTLSCertificateCanBeParsed(t *testing.T) {
	assert := assert.New(t)

	ca_data, key_data := selfSignedCert()
	_, err := tls.X509KeyPair(ca_data, key_data)
	assert.Nil(err)
}

func TestSetTLSSetsTLS(t *testing.T) {
	assert := assert.New(t)

	original_cert, original_key := APITLSData()
	new_cert, new_key := selfSignedCert()
	assert.NotEqual(original_cert, new_cert)
	assert.NotEqual(original_key, new_key)

	err := SetTLS(map[string]string{
		"ca_cert":     string(new_cert),
		"private_key": string(new_key),
	})
	assert.Nil(err)

	current_cert, current_key := APITLSData()
	assert.Equal(current_cert, new_cert)
	assert.Equal(current_key, new_key)
}
