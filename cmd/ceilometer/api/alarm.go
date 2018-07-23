package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/ceilometer/agent"
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
	agent.RegisterAlarm(alarm)
	return nil
}
