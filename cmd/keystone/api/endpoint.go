package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/pkg/domain"
)

type Req struct {
}

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
	req := &Req{}

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
