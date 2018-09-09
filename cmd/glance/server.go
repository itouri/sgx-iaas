package glance

import (
	"github.com/itouri/sgx-iaas/cmd/glance/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/v1/images/:image_id", api.GetImage)
	e.GET("/v1/images/status/:image_id", api.GetImageStatus)
	e.GET("/v1/images/status", api.GetAllImageStatus)
	e.POST("/v1/images", api.PostImage)
	e.DELETE("/v1/images", api.GetImage)

	// TODO
	e.File("/image/", "/home/image/")

	e.Start(":1323")
}
