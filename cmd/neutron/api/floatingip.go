package api

import (
	"github.com/itouri/sgx-iaas/pkg/domain"
	// "github.com/itouri/sgx-iaas/pkg/domain/neutron"
)

func GetAllFloatingIPs(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// json

	return nil
}

func GetFloatingIP(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func PostFloatingIP(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func PutFloatingIP(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func DeleteFloatingIP(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}
