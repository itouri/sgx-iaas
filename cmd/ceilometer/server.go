package ceilometer

import (
	"github.com/itouri/sgx-iaas/cmd/ceilometer/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/v1/alarm", api.PostAlarm())

	e.Start(":1323")
}
