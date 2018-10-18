package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/labstack/echo"
)

var (
	imageCryptoKey string
)

const (
	pubkeyFilePath = "./pub_image_crypto_key.pem"
	prikeyFilePath = "./pri_image_crypto_key.pem"
)

func init() {
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

func GetImageCryptoKey(c echo.Context) error {
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
