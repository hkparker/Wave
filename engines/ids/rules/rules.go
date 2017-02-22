package rules

import (
	"github.com/hkparker/Wave/models"
)

var Alerts = make(chan models.Alert, 0)

func init() {
	// setup js rules if needed
	// forever read alerts, alert and save
}
