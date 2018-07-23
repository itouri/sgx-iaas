package engine

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/itouri/sgx-iaas/pkg/domain/heat"
)

var stacks []heat.Template

func init() {
	stacks = []heat.Template{}
}

func RegisterStack(stack *heat.Template) {
	stacks = append(stacks, *stack)
}

func SendAlarmToCeilometer(tmpl *heat.Template) error {
	// ceiloのAPIを呼ぼう
	//TODO
	url := ""
	bin, err := json.Marshal(tmpl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bin))
	return nil
}

// Alarm.IDをもつAlarmを返す
func GetAlarmWithID(id uuid.uuid) *heat.Alarm {
	for _, alarm := range stacks.Alarms {
		if alarm.ID == id {
			return *alarm
		}
	}
	return nil
}
