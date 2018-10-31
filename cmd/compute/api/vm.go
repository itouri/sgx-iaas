package compute

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/itouri/sgx-iaas/cmd/compute/interactor"
	"github.com/labstack/echo"
)

var (
	vmInteractor    *interactor.VmInteractor
	imageInteractor *interactor.ImageInteractor
	//computeURL   string
	imageStorePath string
	glanceURL      string
)

func init() {
	//vmInteractor = &interactor.VmInteractor{}
	imageStorePath = "./images/"
	vmInteractor = &interactor.VmInteractor{
		Path: imageStorePath,
	}
	glanceURL = "" //TODO
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

	createReqMetadata := c.Param("create_req_metadata")

	filepath := imageStorePath + imageID
	url := glanceURL + "/images/" + imageID

	// Task glanceからimageを取ってくる
	// wgetでいい気もする
	err := imageInteractor.GetFileFromGlance(url, filepath)
	if err != nil {
		fmt.Printf("vm create is failed: %s", err.Error())
		return c.String(http.StatusInternalServerError, "Create VM is failed")
	}

	// convert to uuid
	imageUUID, err := uuid.Parse(imageID)
	if err != nil {
		fmt.Printf("image_id can't convert to UUID: %s", err.Error())
		return c.String(http.StatusInternalServerError, "image_id can't convert to UUID")
	}

	// 起動する
	// unixドメインソケットで master enclave の関数を実行する
	err = vmInteractor.VMCreate(imageUUID, []byte(createReqMetadata))
	if err != nil {
		fmt.Printf("vm create is failed: %s", err.Error())
		return c.String(http.StatusInternalServerError, "Create VM is failed")
	}

	return c.NoContent(http.StatusOK)
}

// func PostVMDelete(c echo.Context) error {

// }

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
