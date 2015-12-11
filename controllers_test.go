package main_test

import (
	. "github.com/hkparker/Wave"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Main", func() {
	It("responds to /", func() {
		router := NewRouter()
		server := httptest.NewServer(router)
		_, err := http.Get("http://" + server.Listener.Addr().String())
		Expect(err).To(BeNil())
		//server.Close()
	})
})
