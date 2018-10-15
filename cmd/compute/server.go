package compute

import (
	"github.com/itouri/sgx-iaas/cmd/compute/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/v1/vm/create/:image_id/", api.PostVMCreate)

	e.Start(":1323")
}
