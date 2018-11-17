package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"

	nonce "github.com/LarryBattle/nonce-golang"
	"github.com/google/uuid"
	"github.com/itouri/sgx-iaas/cmd/sgx/api/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain/ras"
	"github.com/labstack/echo"
)

var (
	imageCryptoKey string
)

var imageMetadataInteractor *interactor.ImageMetadataInteractor
var clientIDInteractor *interactor.ClientIDInteractor
var cryptoInteractor *interactor.CryptoInteractor

const (
	clientIDSize  = 16
	mrenclaveSize = 32
	hmacSize      = 32
	tmpStorePath  = "./tmp-images/"
	glanceURL     = "TODO"
)

func init() {
	imageMetadataInteractor = interactor.NewImageMetadataInteractor()
	clientIDInteractor = interactor.NewClientIDInteractor()
	cryptoInteractor = interactor.NewCryptoInteractor()
}

func PostImage(c echo.Context) error {
	// 本当はクライアントを認証するときにもっと色々な情報が必要なんだと思う
	clientID := c.Param("client_id")
	if clientID != "" {
		return c.String(http.StatusBadRequest, "client_id is not included")
	}

	clientUUID, err := uuid.Parse(clientID)
	if err != nil {
		fmt.Printf("cant convert cliend_id to UUID: %s", err)
		return c.String(http.StatusInternalServerError, "cant convert cliend_id to UUID")
	}

	/* client_idが登録済みか検証 */
	id, err := clientIDInteractor.FindOneByCliendID(clientUUID)
	if id == nil {
		fmt.Printf("cant found UUID: %s", err)
		return c.String(http.StatusInternalServerError, "cant found UUID")
	}

	image, err := c.FormFile("image")
	if err != nil {
		fmt.Println("cant read image: %s", err.Error())
		return c.String(http.StatusInternalServerError, "cant read image")
	}

	/* imageIDを発行する */
	imageUUID := uuid.New()
	if err != nil {
		fmt.Printf("cant generate imageUUID uuid: %s", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = tmpStoreImage(image, imageUUID)
	if err != nil {
		fmt.Println("failed tmpStoreImage :%s", err.Error())
		return c.String(http.StatusBadRequest, "failed tmpStoreImage")
	}

	/* mrenclave を取得する */
	//enclaveSoFilePath :=
	// graphene のやつで計測できないかな... 難しそう
	//mrenclave := "abcdabcdabcdabcdabcdabcdabcdabcd"
	appPath := tmpStorePath + "/get_mrenclave_app"
	soPath := tmpStorePath + "/emain.so"
	out, err := exec.Command(appPath, soPath).Output()

	mrenclave := out
	if len(mrenclave) != mrenclaveSize {
		fmt.Printf("mrenclave size is worng: %d", len(mrenclave))
		return c.String(http.StatusInternalServerError, "mrenclave size is worng")
	}

	// 実行してみて mrenclave を取得する
	// .soを実行して get_mrenclave()を実行する
	// out, err := exec.Command().Output()
	// でてきたmrenclaveをras秘密鍵で復号化 -> いらなくない？

	/* nonceを発行する */
	nonce := nonce.NewToken()

	imd := &ras.ImageMetadata{}
	imageMetadataInteractor.InsertImageMetadata(imd)
	if err != nil {
		fmt.Printf("cant InsertImageMetadata: %s", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	hex, err := clientUUID.MarshalBinary()
	if err != nil {
		fmt.Printf("cant clientUUID.MarshalBinary(): %s", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	/* imageをglanceに登録 */
	// image_metadataの作成
	// imageMetadataSize := clientIDSize + mrenclaveSize + hmacSize

	imageMetadata := []byte{}
	imageMetadata = append(imageMetadata, hex...)
	imageMetadata = append(imageMetadata, []byte(mrenclave)...)

	// hmacを計測(enclaveと同じ方法でやないと)
	hmac := sha256.Sum256(imageMetadata)
	imageMetadata = append(imageMetadata, hmac[:]...)

	// imageMetadata を ras公開鍵で暗号化(enclaveと同じ方法でやないと)
	// cgo使いたい....
	encedIMD, err := cryptoInteractor.Encrypt(imageMetadata)
	if err != nil {
		fmt.Println("failed encrypt imag metada :%s", err.Error())
		return c.String(http.StatusBadRequest, "failed encrypt imag metada")
	}

	// tmp store path はいらないかも
	err = tmpStoreImageMetadata(imageUUID, encedIMD)
	if err != nil {
		fmt.Println("failed tmpStoreImage :%s", err.Error())
		return c.String(http.StatusBadRequest, "failed tmpStoreImage")
	}

	err = imageResgister(imageUUID)
	if err != nil {
		fmt.Println("failed image register :%s", err.Error())
		return c.String(http.StatusBadRequest, "failed image register")
	}

	/* clientに返すもの */
	// image_id, nonce
	ret := imageUUID.String() + "," + nonce
	return c.String(http.StatusOK, ret)
}

func GetClientID(c echo.Context) error {
	/* clientIDを発行する(特に同じクライアントだから同じにする必要もないと思う) -> MRSIGNER で良くない？？ */
	clientID := uuid.New()

	err := clientIDInteractor.InsertClientID(clientID)
	if err != nil {
		fmt.Println("failed InsertClientID: %s\n", err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, clientID.String())
}

func GetCryptoKey(c echo.Context) error {
	return c.File(cryptoInteractor.GetPublicKeyPath())
}

// TODO 同じコードを2回書いてる
func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func tmpStoreImage(image *multipart.FileHeader, imageUUID uuid.UUID) error {
	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	imageID := imageUUID.String()
	// uuid 中のハイフンが邪魔かも
	imageFile, err := os.Create(tmpStorePath + imageID)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	if _, err = io.Copy(imageFile, src); err != nil {
		return err
	}
	return nil
}

// TODO 一々保存しなくて良くない？
// TODO glanceにも同じコード書いてる
func tmpStoreImageMetadata(imageUUID uuid.UUID, imageMetadata []byte) error {
	imageMetadataFile, err := os.Create(tmpStorePath + "metadata")
	if err != nil {
		return err
	}
	defer imageMetadataFile.Close()

	_, err = imageMetadataFile.Write(imageMetadata)
	if err != nil {
		return err
	}
	return nil
}

func imageResgister(imageUUID uuid.UUID) error {
	filename := tmpStorePath + imageUUID.String()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//キーとなる操作
	fileWriter, err := bodyWriter.CreateFormFile("image", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//ファイルハンドル操作をオープンする
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(glanceURL+"/images/"+imageUUID.String(), contentType, bodyBuf)
	if err != nil {
		fmt.Println("Posting feiled: URL:" + glanceURL + "/images")
		return err
	}
	defer resp.Body.Close()
	// resp_body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	return nil
}
