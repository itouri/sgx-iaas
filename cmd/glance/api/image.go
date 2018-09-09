package api

import (
	"io"
	"net/http"
	"os"

	"github.com/itouri/sgx-iaas/cmd/glance/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain"
	uuid "github.com/satori/go.uuid"
)

var imageInteractor *interactor.ImageInteractor

func init() {
	imageInteractor = &interactor.ImageInteractor{
		Path: "/home/image/"
	}
}

func GetImage(c domain.Context) error {
	imageID := c.Param("image_id")

	// imageの参照先のurlを返せばいいのかな

	return nil
}

func GetAllImages(c domain.Context) error {
	req := &Req{}

}

func PostImage(c domain.Context) error {
	file, err := c.FormFile()
	if err != nil {
		return err
	}

	id, err := imageInteractor.StoreFile(file)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, id.String())
}

func DeleteImage(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}
