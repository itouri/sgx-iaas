package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itouri/sgx-iaas/pkg/domain/heat"
	uuid "github.com/satori/go.uuid"
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
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(bin))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	return nil
}

// Alarm.IDをもつAlarmを返す
func GetAlarmWithID(id uuid.UUID) *heat.Alarm {
	for _, stack := range stacks {
		for _, alarm := range stack.Alarms {
			if alarm.ID == id {
				return &alarm
			}
		}
	}
	return nil
}
