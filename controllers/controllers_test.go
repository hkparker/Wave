package controllers

import (
	"net/http"
	"net/http/httptest"
)

var (
	testing_endpoint string
	testing_client   http.Client
)

func init() {
	server := httptest.NewServer(NewRouter())
	testing_endpoint = server.URL
	testing_client = http.Client{}
}
