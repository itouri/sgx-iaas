package image_key_provider

import (
	"github.com/itouri/sgx-iaas/cmd/ceilometer/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ra/image_crypto_key", api.GetImageCryptoKey())

	e.Start(":1323")
}
