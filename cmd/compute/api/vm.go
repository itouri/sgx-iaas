package compute

import (
	"fmt"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/compute/interactor"
	"github.com/labstack/echo"
)

var (
	vmInteractor *interactor.VmInteractor
	computeURL   string
)

func init() {
	vmInteractor = &interactor.VmInteractor{}
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
		return c.String(http.StatusBadRequest, "image_id is lacked")
	}

	// Task glanceからimageを取ってくる

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	return nil
}
