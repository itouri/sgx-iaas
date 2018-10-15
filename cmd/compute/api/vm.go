package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/keystone/interactor"
	"github.com/labstack/echo"
)

var (
	vmInteractor *interactor.VMInteractor
	computeURL   string
)

func init() {
	vmInteractor = &interactor.VMInteractor{}
}

// func GetVMStatus(c echo.Context) error {
// 	VMID := c.Param("vm_id")
// 	if VMID == "" {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}

// 	// Serviceの参照先のurlを返せばいいのかな
// 	service := catalog.GetServiceWithID(serviceID)

// 	ret, err := json.Marshal(service)
// 	if err != nil {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}
// 	return nil
// }

// func GetAllVMStatus(c echo.Context) error {
// 	service := catalog.GetAllServices
// 	ret, err := json.Marshal(service)
// 	if err != nil {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}

// 	return nil
// }

func PostVMCreate(c echo.Context) error {
	imageID := c.Param("image_id")
	if imageID == "" {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO grapheneにどんなことを指示する?

	return nil
}

func DeleteVM(c echo.Context) error {
	imageID := c.Param("image_id")
	if imageID == "" {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO ...

	return nil
}
