package compute

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/itouri/sgx-iaas/cmd/compute/interactor"
	"github.com/labstack/echo"
)

var (
	vmInteractor *interactor.VmInteractor
	//computeURL   string
	imageStorePath string
	glanceURL      string
)

func init() {
	vmInteractor = &interactor.VmInteractor{}
	imageStorePath = "./images/"
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

	filepath := imageStorePath + imageID
	url := glanceURL + "/iamges/" + imageID

	// Task glanceからimageを取ってくる
	// wgetでいい気もする
	if !isExist(filepath) {
		err := exec.Command("wget", url, "-P", filepath).Run()
	}

	// 起動する
	// unixドメインソケットでmaster enclaveの関数を実行する

	return nil
}

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
