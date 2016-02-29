package middleware

import (
	"flag"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	seedMiddlewareTests()
	exitcode := m.Run()
	//teardownUserTests()
	os.Exit(exitcode)
}

func seedMiddlewareTests() {
	helpers.SetEnv("testing")
	helpers.DB().CreateTable(&models.User{})
	user := models.User{
		Name:     "Joe Hackerman",
		Password: []byte{},
		Email:    "joehacker@example.com",
	}
	helpers.DB().Create(&user)
}
