package interactor

import (
	"fmt"
	"io"
	"os"

	"mime/multipart"

	uuid "github.com/satori/go.uuid"
)

type ImageInteractor struct {
	Path string
}

// http.FileServerを使えばいいからいらない
// func (ii *ImageInteractor) GetImage(uuid uuid.UUID)  {
// }

func (ii *ImageInteractor) GetFileStatus() {
	// TODO
}

func (ii *ImageInteractor) GetAllFileStatus() {
	// TODO
}

func (ii *ImageInteractor) StoreFile(file *multipart.FileHeader) (*uuid.UUID, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	id := uuid.Must(uuid.NewV4(), err)
	if err != nil {
		return nil, err
	}

	dstFile, err := os.Create(ii.Path + id.String())
	if err != nil {
		return nil, err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		return nil, err
	}

	return &id, nil
}

// TODO 怖い
func (ii *ImageInteractor) DeleteFile(imageID string) error {
	err := os.Remove("/home/image/" + imageID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
