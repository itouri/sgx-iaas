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

	e.GET("/stacks", api.GetStack)
	e.GET("/stacks", api.GetAllStacks)
	e.POST("/stacks", api.PostStack)

	e.POST("/action/:action_id", api.PostAction)

	e.Start(":1323")
}
