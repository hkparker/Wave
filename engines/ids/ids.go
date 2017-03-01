package ids

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	"github.com/robertkrimen/otto"
)

var VMs = make([]*otto.Otto, 0)
var Alerts = make(chan models.Alert, 0)

func init() {
	go processAlerts()
	buildVMs()
}

var alerting_function = func(call otto.FunctionCall) otto.Value {
	Alerts <- models.Alert{} //call.Argument(0).String()
	return otto.Value{}
}

func processAlerts() {
	for _ = range Alerts {
		// save to database
		// send down websocket
		// update metadata relationships
		// email / message / page
	}
}

func buildVMs() {
	rule_path := "engines/ids/rules"
	rule_files, err := helpers.AssetDir(rule_path)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "ids.buildVMs",
			"error": err.Error(),
		}).Error("unable to load rules")
		return
	}
	for _, rule_file := range rule_files {
		if len(rule_file) < 3 {
			continue
		} else if rule_file[len(rule_file)-3:] != ".js" {
			continue
		}
		if rule_data, ferr := helpers.Asset(rule_path + "/" + rule_file); ferr == nil {
			vm := otto.New()
			_, err := vm.Run(string(rule_data))
			if err != nil {
				log.Error(err)
			}
			vm.Set("alert", alerting_function)
			VMs = append(VMs, vm)
		} else {
			log.WithFields(log.Fields{
				"at":    "ids.buildVMs",
				"error": ferr.Error(),
			}).Error("unable to load rule file")
		}
	}
}

func Insert(frame string, parsed models.Wireless80211Frame) {
	// dedup between collectors
	for _, vm := range VMs {
		_, err := vm.Run(fmt.Sprintf("evaluate(%s)", frame))
		if err != nil {
			log.Error(err)
		}
	}
}
