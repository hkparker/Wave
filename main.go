package main

import (
	"fmt"
	"github.com/hkparker/Wave/controllers"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Setup environment
	helpers.Setup()
	models.Connect()

	// Start Collector server
	go helpers.RunTLS(
		controllers.NewCollector(),
		fmt.Sprintf(
			"%s:%d",
			helpers.WaveBind,
			helpers.CollectorPort,
		),
		models.CollectorTLSConfig(),
	)

	// Start Wave API
	if helpers.TLS {
		helpers.RunTLS(
			controllers.NewAPI(),
			fmt.Sprintf(
				"%s:%d",
				helpers.WaveBind,
				helpers.WavePort,
			),
			models.APITLSConfig(),
		)
	} else {
		controllers.NewAPI().Run(fmt.Sprintf(
			"%s:%d",
			helpers.WaveBind,
			helpers.WavePort,
		))
	}
}
