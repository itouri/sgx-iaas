package api

import (
	"github.com/itouri/sgx-iaas/cmd/heat/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain"
)

var alarmInteractor *interactor.AlarmInteractor

func init() {
	alarmInteractor = &interactor.AlarmInteractor{}
}

type Req struct {
	Template string `json:"template"`
}

func PostAction(c domain.Context) error {
	actionID := c.Param("action_id")
	alarm := interactor.FindByAlarmID(actionID)

}
