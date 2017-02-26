package ids

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func Insert(frame string, parsed models.Wireless80211Frame) {
	log.Info(frame)
}
