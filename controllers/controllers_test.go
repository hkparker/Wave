package controllers

import (
	"github.com/hkparker/Wave/database"
	//"io"
	"net/http/httptest"
)

var (
	//server           *httptest.Server
	//reader           io.Reader
	testing_endpoint string
	testing_client   string
)

func init() {
	server := httptest.NewServer(NewRouter(database.TestDB()))
	testing_endpoint = server.URL
}
