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

	// e.GET("/alarm", api.PostAlarm())

	e.POST("/alarm", api.PostAlarm())
	e.DELETE("/alarm/:alarm_id", api.DeleteAlarm())

	e.Start(":1323")
}
