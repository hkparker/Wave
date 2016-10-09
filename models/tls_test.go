package models

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	//log "github.com/Sirupsen/logrus"
	"testing"
)

func TestDefaultTLSCertificateCanBeParsed(t *testing.T) {
	assert := assert.New(t)

	ca_data, key_data := defaultCA()
	_, err := tls.X509KeyPair(ca_data, key_data)
	assert.Nil(err)
}
