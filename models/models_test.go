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
	//teardownUserTests()
	os.Exit(exitcode)
}

func seedModelsTests() {
	helpers.SetEnv("testing")
	helpers.DB().CreateTable(&User{})
	helpers.DB().CreateTable(&Session{})
	user := User{
		Name:     "Joe Hackerman",
		Password: []byte{},
		Email:    "usertest@example.com",
	}
	helpers.DB().Create(&user)
	user2 := User{
		Name:     "Joe Hackerman",
		Password: []byte{},
		Email:    "joehacker@example.com",
	}
	helpers.DB().Create(&user2)
}
