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

	// e.GET("/v1/alarm", api.PostAlarm())

	e.POST("/v1/alarm", api.PostAlarm())
	e.DELETE("/v1/alarm/:alarm_id", api.DeleteAlarm())

	e.Start(":1323")
}
