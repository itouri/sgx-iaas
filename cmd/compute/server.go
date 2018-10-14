package compute

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.GET("/v1/vm/start/:image_id", GetStartVM)
	//e.GET("/v1/vm/stop/:image_id", GetStopVM)
	e.GET("/v1/vm/create/:image_id", GetCreateVM)

	e.Start(":1323")
}
