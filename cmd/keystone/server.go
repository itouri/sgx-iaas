package keystone

import (
	"github.com/itouri/sgx-iaas/cmd/keystone/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// endpoints
	e.GET("/v1/endpoints", api.GetAllEndPoints)
	e.POST("/v1/endpoints", api.PostEndPoint)

	e.GET("/v1/endpoints/:endpoint_id", api.GetEndPoint)
	e.PATCH("/v1/endpoints/:endpoint_id", api.PatchEndPoint)
	e.DELETE("/v1/endpoints/:endpoint_id", api.DeleteEndPoint)

	// services
	e.GET("/v1/services", api.GetAllServices)
	e.POST("/v1/services", api.PostService)

	e.GET("/v1/services/:service_id", api.GetService)
	e.PATCH("/v1/services/:service_id", api.PatchService)
	e.DELETE("/v1/services/:service_id", api.DeleteService)

	e.Start(":1323")
}
