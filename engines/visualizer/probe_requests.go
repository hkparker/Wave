package visualizer

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func updateProbeRequests(frame models.Wireless80211Frame) {
	if frame.Type == "MgmtProbeReq" {
		//log.Info(string(frame.Elements["SSID"]))
	}
}
