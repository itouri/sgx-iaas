package heat

import (
	"github.com/itouri/sgx-iaas/cmd/heat/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/v1/stacks", api.GetStack)
	e.GET("/v1/stacks", api.GetAllStacks)
	e.POST("/v1/stacks", api.PostImage)

	e.Start(":1323")
}
