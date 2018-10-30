package compute

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.GET("/vm/start/:image_id", GetStartVM)
	//e.GET("/vm/stop/:image_id", GetStopVM)
	e.POST("/vm/create/:image_id", PostVMCreate)

	e.Start(":1323")
}
