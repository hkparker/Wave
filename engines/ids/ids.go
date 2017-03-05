package ids

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/helpers"
	"github.com/hkparker/Wave/models"
	"github.com/robertkrimen/otto"
	"sync"
)

var VMs = make(map[string][]*otto.Otto, 0)
var NewVMs = make(chan []*otto.Otto, 1)
var Alerts = make(chan models.Alert, 0)

func init() {
	go processAlerts()
	go prepareVMs()
}

var alerting_function = func(call otto.FunctionCall) otto.Value {
	new_alert := models.Alert{}
	err := json.Unmarshal([]byte(call.Argument(0).String()), new_alert)
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "ids.alerting_function",
			"error": err.Error(),
		}).Error("bad alert from rule")
	} else {
		Alerts <- new_alert
	}
	return otto.Value{}
}

func processAlerts() {
	for _ = range Alerts {
		// dedup between interfaces
		// save to database
		// send down websocket
		// update metadata relationships
		// email / message / page
	}
}

func prepareVMs() {
	for {
		NewVMs <- buildVMs()
	}
}

func buildVMs() (vm_set []*otto.Otto) {
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
				log.WithFields(log.Fields{
					"at":    "ids.buildVMs",
					"file":  rule_file,
					"error": err.Error(),
				}).Error("error loading rule data into VM")
			}
			vm.Set("alert", alerting_function)
			vm_set = append(vm_set, vm)
		} else {
			log.WithFields(log.Fields{
				"at":    "ids.buildVMs",
				"error": ferr.Error(),
			}).Error("unable to load rule file")
		}
	}
	return
}

func Insert(frame string, parsed models.Wireless80211Frame) {
	vm_set, ok := VMs[parsed.Interface]
	if !ok {
		vm_set = <-NewVMs
		VMs[parsed.Interface] = vm_set
	}
	var evals sync.WaitGroup
	for _, vm := range vm_set {
		evals.Add(1)
		go func(vm *otto.Otto) {
			_, err := vm.Run(fmt.Sprintf("evaluate(%s)", frame))
			if err != nil {
				log.WithFields(log.Fields{}).Error(err)
			}
		}(vm)
	}
	evals.Wait()
}
