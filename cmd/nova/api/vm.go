package api

import (
	"encoding/json"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/keystone/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain"
)

var vmInteractor *interactor.VMInteractor

func init() {
	vmInteractor = &interactor.VMInteractor{}
}

func GetVMStatus(c domain.Context) error {
	VMID := c.Param("vm_id")
	if VMID == "" {
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

func GetAllVMStatus(c domain.Context) error {
	service := catalog.GetAllServices
	ret, err := json.Marshal(service)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func PostVMCreate(c domain.Context) error {
	imageID := c.Param("image_id")
	if imageID == "" {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO scheduling

	// 
	

	return nil
}

func DeleteService(c domain.Context) error {
	imageID := c.Param("image_id")
	if imageID == "" {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO ...


	return nil
}
