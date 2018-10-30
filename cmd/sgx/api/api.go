package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	nonce "github.com/LarryBattle/nonce-golang"
	"github.com/itouri/sgx-iaas/cmd/sgx/api/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain/ras"
	"github.com/labstack/echo"
	"github.com/google/uuid"
)

var (
	imageCryptoKey string
)

var imageMetadataInteractor *interactor.ImageMetadataInteractor
var clientIDInteractor *interactor.ClientIDInteractor

const (
	pubkeyFilePath = "./pub_image_crypto_key.pem"
	prikeyFilePath = "./pri_image_crypto_key.pem"
)

func init() {
	imageMetadataInteractor = interactor.NewImageMetadataInteractor()
	clientIDInteractor = interactor.NewClientIDInteractor()

	if !isExist(pubkeyFilePath) {
		err := generateKey()
		if err != nil {
			fmt.Printf("Cant generate image_crypto_key:" + err.Error())
			return
		}
	}

	file, err := os.Open(pubkeyFilePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("Cant read ./image_crypto_key.pub")
		return
	}

	buf := make([]byte, 1024)
	for {
		// TODO 合ってる?
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		imageCryptoKey += string(buf)
	}
}

func PostImage(c echo.Context) error {
	// 本当はクライアントを認証するときにもっと色々な情報が必要なんだと思う
	clientIDstr := c.Param("client_id")
	if clientIDstr != "" {
		return c.String(http.StatusBadRequest, "client_id is not included")
	}

	clientID, err := uuid.FromString(clientIDstr)
	if err != nil {
		fmt.Printf("cant convert cliend_id to UUID: %s", err)
		return c.String(http.StatusInternalServerError, "cant convert cliend_id to UUID")
	}

	/* client_idが登録済みか検証 */
	id, err := clientIDInteractor.FindOneByCliendID(clientID)
	if id == nil {
		fmt.Printf("cant found UUID: %s", err)
		return c.String(http.StatusInternalServerError, "cant found UUID")
	}

	/* mrenclave を取得する */
	enclaveSoFilePath := 

	// 実行してみて mrenclave を取得する
	// .soを実行して get_mrenclave()を実行する
	out, err := exec.Command().Output()

	/* imageIDを発行する */
	imageID := uuid.New()
	if err != nil {
		fmt.Printf("cant generate imageID uuid: %s", err.Error())
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

	/* clientに返すもの */
	// image_id, nonce
	ret := imageID.String() + "," + nonce
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
	return c.File(pubkeyFilePath)
}

// TODO 同じコードを2回書いてる
func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func generateKey() error {
	// TODO できればECが良い
	// copy from https://gist.github.com/sdorra/1c95de8cb80da31610d2ad767cd6f251
	// gob?
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return err
	}
	publicKey := key.PublicKey

	err = savePEMKey(prikeyFilePath, key)
	if err != nil {
		return err
	}

	err = savePublicPEMKey(pubkeyFilePath, publicKey)
	if err != nil {
		return err
	}

	return nil
}

func savePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		return err
	}

	return nil
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		return err
	}

	return nil
}
