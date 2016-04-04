package models

import (
	"flag"
	"github.com/hkparker/Wave/helpers"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	seedModelsTests()
	exitcode := m.Run()
	teardownModelsTests()
	os.Exit(exitcode)
}

func seedModelsTests() {
	helpers.SetEnv("testing")
	helpers.DB().CreateTable(&User{})
	helpers.DB().CreateTable(&Session{})
	user := User{
		Name:  "Joe Hackerman",
		Email: "modeltest1@example.com",
	}
	helpers.DB().Create(&user)
}

func teardownModelsTests() {
	helpers.DB().Unscoped().Where(
		"email LIKE ?",
		"modeltest%",
	).Delete(User{})
}
