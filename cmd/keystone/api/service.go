package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/pkg/domain"
)

func GetService(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Serviceの参照先のurlを返せばいいのかな

	return nil
}

func GetAllServices(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func PostService(c domain.Context) error {
	type Req struct {
		Name   string `json:"name" validate:"required"`
		Type   string `json:"type" validate:"required"`
		IPAddr string `json:"ip_address" validate:"required"`
	}
	req := new(Req)

	if err := c.BindValidate(u); err != nil {
		return err
	}

	return nil
}

func DeleteService(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}
