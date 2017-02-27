package ids

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hkparker/Wave/models"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"os"
)

var VMs = make([]*otto.Otto, 0)

func init() {
	//loadRules()
	buildVMs()
}

func buildVMs() {
	rule_path := "engines/ids/rules/"
	infos, err := ioutil.ReadDir(rule_path)
	if err != nil {
		log.Error(err.Error())
		return
	}
	for _, info := range infos {
		if info.IsDir() {
			manifest_file := rule_path + info.Name() + "/manifest.json"
			rule_file := rule_path + info.Name() + "/rule.js"
			if _, err := os.Stat(manifest_file); os.IsNotExist(err) {
				log.Warn(manifest_file)
				continue
			}
			if _, err := os.Stat(rule_file); os.IsNotExist(err) {
				log.Warn(rule_file)
				continue
			}
			rule_func_data, err := ioutil.ReadFile(rule_file)
			if err != nil {
				log.Error(err)
				continue
			}
			vm := otto.New()
			vm.Run(string(rule_func_data))
			//vm.Set("alert", func(call otto.FunctionCall) otto.Value {
			//	alerts <- call.Argument(0).String()
			//	return otto.Value{}
			//})
			VMs = append(VMs, vm)
		}
	}
}

func Insert(frame string, parsed models.Wireless80211Frame) {
	for _, vm := range VMs {
		vm.Run(fmt.Sprintf("evaluate(%s)", frame))
	}
}
