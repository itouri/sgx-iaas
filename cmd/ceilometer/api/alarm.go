package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/pkg/domain"
)

type Req struct {
}

// alarm structを受け取る
func PostAlarm(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}
