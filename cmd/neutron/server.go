package neutron

import (
	"github.com/itouri/sgx-iaas/neutron/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// floatingips
	e.GET("/v1/floatingips", api.GetAllFloatingIPs)
	e.GET("/v1/floatingips/:floatingip_id", api.GetFloatingIP)
	e.POST("/v1/floatingips", api.PostFloatingIP)
	e.PUT("/v1/floatingips/:floatingip_id", api.PutFloatingIP)
	e.DELETE("/v1/floatingips/:floatingip_id", api.DeleteFloatingIP)

	// network
	e.GET("/v1/networks", api.GetAllNetworks)
	e.GET("/v1/networks/:network_id", api.GetNetwork)
	e.POST("/v1/networks", api.PostNetwork)
	e.PUT("/v1/networks/:network_id", api.PutNetwork)
	e.DELETE("/v1/networks/:network_id", api.DeleteNetwork)

	// router
	e.GET("/v1/routers", api.GetAllRouters)
	e.GET("/v1/routers/:router_id", api.GetRouter)
	e.POST("/v1/routers", api.PostRouter)
	e.PUT("/v1/routers/:router_id", api.PutRouter)
	e.DELETE("/v1/routers/:router_id", api.DeleteRouter)

	e.GET("/v1/stacks", api.GetStack)
	e.GET("/v1/stacks", api.GetAllStacks)
	e.POST("/v1/stacks", api.PostImage)

	e.Start(":1323")
}
