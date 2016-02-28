package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"testing"
)

var db = seedAuthenticationTests()

func seedAuthenticationTests() gorm.DB {
	// drop test database if it exists
	db, err := gorm.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to connect to postgres")
	}
	db.CreateTable(&models.User{})
	user := models.User{
		Name:     "Joe Hackerman",
		Password: []byte{},
		Email:    "joehacker@example.com",
	}
	db.Create(&user)
	return db
}

//func TestPublicResourcesUnrestricted(t *testing.T) {
//	user := models.User{}
//	db.First(&user)
//	if user.Name != "Joe Hackerman" {
//		t.Fatal("postgres test db didn't seed")
//	}
//}

func TestActiveSessionTrueWithSessionAndUpdatesTime(t *testing.T) {

}
