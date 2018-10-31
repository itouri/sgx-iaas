package interactor

import (
	"fmt"
	"io"
	"os"

	"mime/multipart"

	"github.com/google/uuid"
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

func (ii *ImageInteractor) StoreFile(file *multipart.FileHeader, imageUUID uuid.UUID) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	imageID := imageUUID.String()
	// uuid 中のハイフンが邪魔かも
	dstFile, err := os.Create(ii.Path + imageID)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		return err
	}

	return nil
}

// TODO 怖い セキュリティ的にガバガバ
func (ii *ImageInteractor) DeleteFile(imageID string) error {
	filepath := ii.Path + imageID

	if !isExist(filepath) {
		return fmt.Errorf("image is not found:" + imageID)
	}

	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
