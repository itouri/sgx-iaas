package compute

import (
<<<<<<< HEAD
	"github.com/itouri/sgx-iaas/cmd/compute/api"
=======
>>>>>>> d94fada8e0371209689bca177dc55f4ad60cd05d
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

<<<<<<< HEAD
	e.POST("/v1/vm/create/:image_id/", api.PostVMCreate)
=======
	//e.GET("/v1/vm/start/:image_id", GetStartVM)
	//e.GET("/v1/vm/stop/:image_id", GetStopVM)
	e.GET("/v1/vm/create/:image_id", GetCreateVM)
>>>>>>> d94fada8e0371209689bca177dc55f4ad60cd05d

	e.Start(":1323")
}
