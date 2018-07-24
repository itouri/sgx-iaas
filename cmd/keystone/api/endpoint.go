package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/pkg/domain"
)

func GetEndpoint(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func GetAllEndpoints(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func PostEndpoint(c domain.Context) error {
	type Req struct {
		Enabled   string `json:enabled`
		ServiceID string `json:service_id`
		URL       string `json:url`
	}
	req := new(Req)
	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return nil
}

func DeleteEndpoint(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}
