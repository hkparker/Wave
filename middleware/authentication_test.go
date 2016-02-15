package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"testing"
)

func seedAuthenticationTests() gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to postgres")
	}
	return db
}

func TestPublicResourcesUnrestricted(t *testing.T) {
	seedAuthenticationTests()
}
