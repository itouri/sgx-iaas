package api

import (
	"encoding/json"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/keystone/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
)

var serviceInteractor *interactor.ServiceInteractor

func init() {
	serviceInteractor = &interactor.ServiceInteractor{}
}

func GetServiceResolve(c domain.Context) error {
	serviceTypeStr := c.Param("service_type")
	serviceType := keystone.ToEnumServiceType(serviceTypeStr)
	if serviceType == -1 {
		return err
	}

	service, err := serviceInteractor.FindByServiceType(serviceType)
	if err := nil {
		return err
	}

	res, err := json.Marshal(service)
	if err := nil {
		return err
	}
	
	return c.JSON(http.StatusOK, res)
}

func GetService(c domain.Context) error {
	serviceID := c.Param("service_id")
	if serviceID == "" {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Serviceの参照先のurlを返せばいいのかな
	service := catalog.GetServiceWithID(serviceID)
	ret, err := json.Marshal(service)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetAllServices(c domain.Context) error {
	service := catalog.GetAllServices
	ret, err := json.Marshal(service)
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
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func DeleteService(c domain.Context) error {
	serviceID := c.Param("service_id")

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}
