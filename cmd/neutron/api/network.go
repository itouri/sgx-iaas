package api

import (
	"github.com/itouri/sgx-iaas/pkg/domain"
	// "github.com/itouri/sgx-iaas/pkg/domain/neutron"
)

func GetAllNetworks(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func GetNetwork(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func PostNetwork(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func PutNetwork(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func DeleteNetwork(c domain.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}
