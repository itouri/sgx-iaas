package main

import (
	"fmt"

	"github.com/itouri/sgx-iaas/cmd/keystone/util"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := util.RegisterEndpoint(keystone.RAKey, "localhost", 22222)
	if err != nil {
		fmt.Println("Error RegisterEndpoint: " + err.Error())
		return
	}

	e.GET("/ra/image_crypto_key", GetImageCryptoKey)

	e.Start(":22222")
}
