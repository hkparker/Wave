package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	testing_endpoint string
	testing_client   http.Client
)

func init() {
	os.Setenv("WAVE_ENV", "testing")
	server := httptest.NewServer(NewRouter())
	testing_endpoint = server.URL
	testing_client = http.Client{}
}
