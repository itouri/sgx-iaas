package main

import (
	"github.com/itouri/sgx-iaas/cmd/keystone/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// endpoint -> serviceに統一
	// e.GET("/endpoints", api.GetAllEndPoints)
	// e.POST("/endpoints", api.PostEndPoint)

	// e.GET("/endpoints/:endpoint_id", api.GetEndPoint)
	// e.PATCH("/endpoints/:endpoint_id", api.PatchEndPoint)
	// e.DELETE("/endpoints/:endpoint_id", api.DeleteEndPoint)

	// services
	//e.GET("/services", api.GetAllServices)
	e.GET("/services/resolve/:service_type", api.GetServiceResolve)
	e.POST("/services", api.PostService)

	//e.GET("/services/:service_id", api.GetService)
	//e.PATCH("/services/:service_id", api.PatchService)
	e.DELETE("/services/:service_id", api.DeleteService)

	e.Start(":1323")
}
