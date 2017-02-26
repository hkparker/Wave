package helpers

import (
	"crypto/tls"
	log "github.com/Sirupsen/logrus"
	"net"
	"net/http"
)

func RunTLS(handler http.Handler, address string, config *tls.Config) {
	server := &http.Server{
		Handler: handler,
	}
	tcp_listener, err := net.Listen("tcp", address)
	if err != nil {
		log.WithFields(log.Fields{
			"address": address,
			"error":   err.Error(),
		}).Fatal("unable to create tls listener")
	}
	tls_listener := tls.NewListener(tcp_listener, config)
	server.Serve(tls_listener)
}
