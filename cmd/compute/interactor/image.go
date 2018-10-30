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

func (ii *ImageInteractor) StoreFile(file *multipart.FileHeader) (*uuid.UUID, error) {
	src, err := file.Open()
	if err != nil {
		fmt.Println("file.Open()")
		return nil, err
	}
	defer src.Close()

	id := uuid.New()
	if err != nil {
		fmt.Println("uuid.New()")
		return nil, err
	}

	// uuid 中のハイフンが邪魔かも
	dstFile, err := os.Create(ii.Path + id.String())
	if err != nil {
		fmt.Println("os.Create(ii.Path + id.String())")
		return nil, err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		fmt.Println("io.Copy")
		return nil, err
	}

	return &id, nil
}

// TODO 怖い セキュリティ的にガバガバ
func (ii *ImageInteractor) DeleteFile(imageID string) error {
	filepath := ii.Path + "/" + imageID

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
