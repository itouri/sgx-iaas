package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/keystone/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/labstack/echo"
)

var serviceInteractor *interactor.ServiceInteractor

func init() {
	serviceInteractor = interactor.NewServiceInteractor()
}

func GetServiceResolve(c echo.Context) error {
	serviceTypeStr := c.Param("service_type")
	serviceType := keystone.ToEnumServiceType(serviceTypeStr)

	service, err := serviceInteractor.FindByServiceType(serviceType)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	type Res struct {
		/// Name   string `json:"name"`
		Type   int    `json:"type"`
		Port   int    `json:"port"`
		IPAddr string `json:"ipaddr"`
	}

	resp := &Res{
		Type:   int(service.Type),
		Port:   int(service.Port),
		IPAddr: service.IPAddr,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf(string(res))

	// TODO なぜBASE64エンコードされている？
	return c.JSON(http.StatusOK, res)
}

// func GetService(c echo.Context) error {
// 	serviceID := c.Param("service_id")
// 	if serviceID == "" {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}

// 	// Serviceの参照先のurlを返せばいいのかな
// 	service := catalog.GetServiceWithID(serviceID)
// 	ret, err := json.Marshal(service)
// 	if err != nil {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}
// 	return nil
// }

// func GetAllServices(c echo.Context) error {
// 	service := catalog.GetAllServices
// 	ret, err := json.Marshal(service)
// 	if err != nil {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}

// 	return nil
// }

func PostService(c echo.Context) error {
	type Req struct {
		/// Name   string `json:"name" validate:"required"`
		Type   string `bson:"type" json:"type" validate:"required"`
		Port   int    `bson:"port" json:"port" validate:"required"`
		IPAddr string `bson:"ipaddr" json:"ipaddr" validate:"required"`
	}
	req := new(Req)

	// TODO validate
	err := c.Bind(req)
	fmt.Println(req)
	if err != nil {
		fmt.Println("Error: c.Bind(req)")
		return c.String(http.StatusBadRequest, err.Error())
	}

	port := uint64(req.Port)

	//ipaddr := net.ParseIP(req.IPAddr)

	service := keystone.Service{
		Type:   keystone.ToEnumServiceType(req.Type), // TODO これ変な数字来たらどうなんの
		Port:   port,
		IPAddr: req.IPAddr, // net.IPにキャストしなくていいの？
	}

	// TODO もし同じtype, port, ipaddrがあったら再起動
	// 現状typeが同じ場合上書きする
	err = serviceInteractor.UpsertService(service)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteService(c echo.Context) error {
	serviceID := c.Param("service_id")
	if serviceID != "" {
		return c.String(http.StatusBadRequest, "please include service_id")
	}

	err := serviceInteractor.DeleteService(serviceID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}
