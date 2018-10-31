package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"

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
)

func init() {
	imageMetadataInteractor = interactor.NewImageMetadataInteractor()
	clientIDInteractor = interactor.NewClientIDInteractor()
	cryptoInteractor = interactor.CryptoInteractor()
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

	/* mrenclave を取得する */
	//enclaveSoFilePath :=
	mrenclave := "abcdabcdabcdabcdabcdabcdabcdabcd"

	// 実行してみて mrenclave を取得する
	// .soを実行して get_mrenclave()を実行する
	// out, err := exec.Command().Output()
	// でてきたmrenclaveをras秘密鍵で復号化

	/* imageIDを発行する */
	imageUUID := uuid.New()
	if err != nil {
		fmt.Printf("cant generate imageUUID uuid: %s", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

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
	// encedIMD := encrypt(imageMetadata, pubkey)

	/* clientに返すもの */
	// image_id, nonce
	ret := imageUUID.String() + "," + nonce
	return c.String(http.StatusOK, ret)
}

func GetClientID(c echo.Context) error {
	var err error
	/* clientIDを発行する(特に同じクライアントだから同じにする必要もないと思う) -> MRSIGNER で良くない？？ */
	clientID := uuid.New()
	if err != nil {
		fmt.Println("uuid.New()")
		return c.String(http.StatusBadRequest, err.Error())
	}

	// DBにclient_idを登録
	clientIDstr := clientID.String()

	return c.String(http.StatusOK, clientIDstr)
}

func GetCryptoKey(c echo.Context) error {
	return c.File(cryptoInteractor.GetPublicKeyPath())
}

// TODO 同じコードを2回書いてる
func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
