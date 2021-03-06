package interactor

import (
	"fmt"
	"io"
	"os"
	"os/exec"

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

func (ii *ImageInteractor) GetImageMetedata(imageID string) ([]byte, error) {
	file, err := os.Open(ii.Path + imageID)
	if err != nil {
		// Openエラー処理
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 1024)
	size := 0
	for {
		size, err := file.Read(buf)
		if size == 0 {
			break
		}
		if err != nil {
			// Readエラー処理
			return nil, err
		}
	}
	return buf[:size], nil
}

func (ii *ImageInteractor) GetFileFromGlance(url string, imageID string) error {
	imagePath := ii.Path + imageID
	if !isExist(imagePath) {
		err := exec.Command("wget", url, "-P", ii.Path).Run()
		if err != nil {
			return err
		}
	}
	return nil
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
