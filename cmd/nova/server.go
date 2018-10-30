package nova

import (
	"github.com/itouri/sgx-iaas/cmd/keystone/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/vm/status/:image_id", api.GetAllServices)
	e.GET("/vm/status/", api.GetAllServices)
	e.POST("/vm/:image_id/create", api.PostService)

	e.Start(":1323")
}
