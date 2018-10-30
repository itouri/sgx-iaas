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
	e.GET("/floatingips", api.GetAllFloatingIPs)
	e.GET("/floatingips/:floatingip_id", api.GetFloatingIP)
	e.POST("/floatingips", api.PostFloatingIP)
	e.PUT("/floatingips/:floatingip_id", api.PutFloatingIP)
	e.DELETE("/floatingips/:floatingip_id", api.DeleteFloatingIP)

	// network
	e.GET("/networks", api.GetAllNetworks)
	e.GET("/networks/:network_id", api.GetNetwork)
	e.POST("/networks", api.PostNetwork)
	e.PUT("/networks/:network_id", api.PutNetwork)
	e.DELETE("/networks/:network_id", api.DeleteNetwork)

	// router
	e.GET("/routers", api.GetAllRouters)
	e.GET("/routers/:router_id", api.GetRouter)
	e.POST("/routers", api.PostRouter)
	e.PUT("/routers/:router_id", api.PutRouter)
	e.DELETE("/routers/:router_id", api.DeleteRouter)

	e.GET("/stacks", api.GetStack)
	e.GET("/stacks", api.GetAllStacks)
	e.POST("/stacks", api.PostImage)

	e.Start(":1323")
}
