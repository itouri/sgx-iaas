package api

import (
	"io"
	"net/http"
	"os"

	"github.com/itouri/sgx-iaas/pkg/domain"
	uuid "github.com/satori/go.uuid"
)

const imagePath = "./images/"

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

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	id := uuid.Must(uuid.NewV4())

	dstFile, err := os.Create(imagePath + id)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, id)
}

func DeleteImage(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return nil
}
