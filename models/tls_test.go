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
