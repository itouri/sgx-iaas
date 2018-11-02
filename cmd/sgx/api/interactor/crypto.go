package interactor

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	publicKeyPath  = "key/RAS-rsa-key-pub.pem"
	privateKeyPath = "key/RAS-rsa-key.pem"
)

type CryptoInteractor struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

var cryptoInteractor *CryptoInteractor

// http://increment.hatenablog.com/entry/2017/08/25/223915
func readRsaPrivateKey() (*rsa.PrivateKey, error) {
	bytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("invalid private key data")
	}

	var key *rsa.PrivateKey
	if block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("key type is not RSA PRIVATE KEY: %s", block.Type)
	}

	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	key.Precompute()

	if err := key.Validate(); err != nil {
		return nil, err
	}
	return key, nil
}

func readRsaPublicKey() (*rsa.PublicKey, error) {
	bytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("invalid public key data")
	}
	if block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("invalid public key type : %s", block.Type)
	}

	// type これで合ってる？
	keyInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// key, ok := keyInterface.(*rsa.PublicKey)
	// if !ok {
	// 	return nil, errors.New("not RSA public key")
	// }

	return keyInterface, nil
}

func NewCryptoInteractor() *CryptoInteractor {
	privateKey, err := readRsaPrivateKey()
	if err != nil {
		panic("failed readRsaPrivateKey: " + err.Error())
	}

	publicKey, err := readRsaPublicKey()
	if err != nil {
		panic("failed readRsaPublicKey: " + err.Error())
	}

	if cryptoInteractor == nil {
		cryptoInteractor = &CryptoInteractor{
			privateKey: privateKey,
			publicKey:  publicKey,
		}
	}
	return cryptoInteractor
}

func (ci *CryptoInteractor) Encrypt(data []byte) ([]byte, error) {
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, ci.publicKey, data)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return nil, err
	}
	return cipherText, nil
}

func (ci *CryptoInteractor) Decrypt(data []byte) ([]byte, error) {
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, ci.privateKey, data)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return nil, err
	}
	return decryptedText, nil
}

func (ci *CryptoInteractor) GetPublicKeyPath() string {
	return publicKeyPath
}
