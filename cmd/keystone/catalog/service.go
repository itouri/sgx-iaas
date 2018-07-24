package catalog

import (
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/satori/go.uuid"
)

var services []keystone.Service

func init() {
	services = []keystone.Service{}
}

// func RegisterEndpointToService(ep *keystone.EndPoint, svid uuid.UUID) {
// 	for _, s := range services {
// 		if s.ID == svid {
// 			s.EndPoints = append(s.EndPoints, *ep)
// 		}
// 	}
// }

func RegisterService(service *keystone.Service) {
	services = append(services, *service)
}

func DeleteService(id uuid.UUID) {
	delIndex := -1
	for i, s := range services {
		if s.ID == id {
			delIndex = i
			break
		}
	}
	if delIndex == -1 {
		return
	}
	services = append(services[:delIndex], services[delIndex+1:]...)
}
