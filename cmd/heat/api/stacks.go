package api

import (
	"fmt"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/heat/engine"
	"github.com/itouri/sgx-iaas/pkg/domain"
	"github.com/itouri/sgx-iaas/pkg/domain/heat"

	uuid "github.com/satori/go.uuid"
	yaml "gopkg.in/yaml.v2"
)

type Req struct {
	Template string `json:"template"`
}

func GetStacks(c domain.Context) error {

}

func PostStack(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	template := &heat.Template{}

	// yamlファイルの解釈
	err = yaml.Unmarshal([]byte(req.Template), template)
	if err != nil {
		//TODO StatusOKではない
		return c.String(http.StatusOK, err.Error())
	}

	alarmMap := map[string]uuid.UUID

	//TODO これだとtemplateに反映されてない?
	// AlarmにUUIDを割り当てる
	for _, alarm := range template.Alarms {
		alarm.ID = uuid.Must(uuid.NewV4())
		alarmMap[alarm.AlarmAction] = alarm.ID
	}

	for _, sp := range template.ScalingPolicies {
		id, ok := alarmMap[sp.Name]
		if ok {
			sp.AlarmID = id
		} else {
			return fmt.Errorf("doesn't exist scaling policis maped alarm ID")
		}
	}

	//TODO valitation
	engine.RegisterStack(template)

	return nil
}
