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
	teardownMiddlewareTests()
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

func teardownMiddlewareTests() {
	helpers.TestDB().Unscoped().Where(
		"email LIKE ?",
		"middlewaretest%",
	).Delete(models.User{})
}

func TestAuthSeeded(t *testing.T) {
	user := models.User{}
	helpers.TestDB().First(&user)
	assert.Equal(t, "Joe Hackerman", user.Name)
}
