package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/ceilometer/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain"
	"github.com/itouri/sgx-iaas/pkg/domain/heat"
)

// alarm structを受け取る
func PostAlarm(c domain.Context) error {
	alarm := new(heat.Alarm)
	err := c.Bind(alarm)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	interactor.InsertAlarm(alarm)
	return nil
}

func DeleteAlarm(c domain.Context) error {
	alarmID := c.Param("alarm_id")
	interactor.DeleteAlarm(alarmID)
	return nil
}
