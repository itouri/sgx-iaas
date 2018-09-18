package api

import (
	"fmt"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/glance/interactor"
	"github.com/labstack/echo"
)

var imageInteractor *interactor.ImageInteractor

func init() {
	imageInteractor = &interactor.ImageInteractor{
		Path: "/home/image/",
	}
}

// e.File("/image/", "/home/image/")が代用
// func GetImage(c echo.Context) error {
// 	imageID := c.Param("image_id")

// 	// imageの参照先のurlを返せばいいのかな

// 	return nil
// }

// func GetAllImages(c echo.Context) error {
// 	req := &Req{}

// }

func PostImage(c echo.Context) error {

	file, err := c.FormFile("image")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	id, err := imageInteractor.StoreFile(file)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return c.String(http.StatusOK, id.String())
}

func DeleteImage(c echo.Context) error {
	imageID := c.Param("image_id")
	if imageID == "" {
		return c.String(http.StatusBadRequest, "image_id is not included")
	}
	imageInteractor.DeleteFile(imageID)

	return nil
}
