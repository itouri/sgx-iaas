package api

import (
	"github.com/itouri/sgx-iaas/cmd/heat/engine"
	"github.com/itouri/sgx-iaas/pkg/domain"
)

type Req struct {
	Template string `json:"template"`
}

func PostAction(c domain.Context) error {
	actionID := c.Param("action_id")
	alarm := engine.GetAlarmWithID(actionID)

}
