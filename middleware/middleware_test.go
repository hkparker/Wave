package middleware

import (
	"flag"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	seedMiddlewareTests()
	exitcode := m.Run()
	os.Exit(exitcode)
}

func seedMiddlewareTests() {
	//helpers.SetEnv("testing")
	// helpers.TestDB().CreateTable(&models.User{})
	user := models.User{
		Name:  "Joe Hackerman",
		Email: "middlewaretest1@example.com",
	}
	helpers.TestDB().Create(&user)
}

func skipTestAuthSeeded(t *testing.T) {
	user := models.User{}
	helpers.TestDB().First(&user)
	assert.Equal(t, "Joe Hackerman", user.Name)
}
