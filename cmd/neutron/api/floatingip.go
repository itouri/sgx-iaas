package api

import (
	"github.com/itouri/sgx-iaas/pkg/domain"
	// "github.com/itouri/sgx-iaas/pkg/domain/neutron"
)

func GetAllFloatingIPs(c echo.Context) error {
	endpointID := c.Param("endpoint_id")

	// json

	return nil
}

func GetFloatingIP(c echo.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func PostFloatingIP(c echo.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func PutFloatingIP(c echo.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}

func DeleteFloatingIP(c echo.Context) error {
	endpointID := c.Param("endpoint_id")

	// Endpointの参照先のurlを返せばいいのかな

	return nil
}
