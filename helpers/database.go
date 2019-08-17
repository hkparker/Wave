package helpers

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var DBHostname string
var DBUsername string
var DBPassword string
var DBAdminUsername string
var DBAdminPassword string
var DBName string
var DBTLS string

func setDatabase() {
	DBHostname = os.Getenv("WAVE_DB_HOSTNAME")
	if DBHostname == "" {
		log.Fatal("WAVE_DB_HOSTNAME envar must be provided")
	}

	DBUsername = os.Getenv("WAVE_DB_USERNAME")
	DBAdminUsername = os.Getenv("WAVE_DB_ADMIN_USERNAME")
	if DBUsername == "" {
		log.Fatal("WAVE_DB_USERNAME envar must be provided")
	}

	DBPassword = os.Getenv("WAVE_DB_PASSWORD")
	DBAdminPassword = os.Getenv("WAVE_DB_ADMIN_PASSWORD")
	if DBPassword == "" {
		log.Fatal("WAVE_DB_PASSWORD envar must be provided")
	}

	DBName = os.Getenv("WAVE_DB_NAME")
	if DBName == "" {
		log.Fatal("WAVE_DB_NAME envar must be provided")
	}

	DBTLS = os.Getenv("WAVE_DB_TLS")
	if DBTLS == "" {
		log.Fatal("WAVE_DB_TLS envar must be provided")
	}
}
