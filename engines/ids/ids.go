package ids

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
)

func Insert(frame models.Wireless80211Frame) {
	str, _ := json.Marshal(frame)
	log.Info(string(str))
}
