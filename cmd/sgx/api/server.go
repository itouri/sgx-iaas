package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// err := util.RegisterEndpoint(keystone.RAKey, "localhost", 22222)
	// if err != nil {
	// 	fmt.Println("Error RegisterEndpoint: " + err.Error())
	// 	return
	// }

	//e.GET("/ra/image_crypto_key", GetImageCryptoKey)
	e.POST("/ra/images", PostImage)

	e.Start(":22222")
}
