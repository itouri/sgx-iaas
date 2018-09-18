package image_key_provider

import (
	"fmt"
	"net/http"
	"os"

	"github.com/itouri/sgx-iaas/pkg/domain"
)

var imageCryptoKey string

func init() {
	file, err := os.Open("./image_crypto_key.pub")
	defer file.Close()
	if err != nil {
		fmt.Printf("Cant read ./image_crypto_key.pub")
		return
	}

	buf := make([]byte, BUFSIZE)
	for {
		// TODO 合ってる?
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Printf(err.String())
			return
		}
		imageCryptoKey += buf
	}
}

func GetImageCryptoKey(c echo.Context) error {
	return c.String(http.StatusOK, imageCryptoKey)
}
