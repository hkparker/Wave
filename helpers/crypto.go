package helpers

import (
	"crypto/rand"
	log "github.com/Sirupsen/logrus"
	"github.com/jbenet/go-base58"
)

func RandomString() (random_string string) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "helpers.RandomString",
			"error": err.Error(),
		}).Fatal("error generating cryptographically secure random string")
	}
	random_string = base58.Encode(bytes)
	return
}
