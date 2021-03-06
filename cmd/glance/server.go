package main

import (
	"fmt"

	"github.com/itouri/sgx-iaas/cmd/keystone/util"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"

	"github.com/itouri/sgx-iaas/cmd/glance/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := util.RegisterEndpoint(keystone.Glance, "localhost", 11111)
	if err != nil {
		fmt.Println("Error RegisterEndpoint: " + err.Error())
		return
	}

	e.GET("/images/:image_id", api.GetImage)
	//e.GET("/images/status/:image_id", api.GetImageStatus)
	//e.GET("/images/status", api.GetAllImageStatus)
	e.POST("/images", api.PostImage)
	e.DELETE("/images/:image_id", api.DeleteImage)

	// TODO
	//e.File("/image", "./image/")

	e.Start(":11111")
}
