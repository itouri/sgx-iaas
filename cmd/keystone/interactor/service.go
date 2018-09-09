package interactor

import (
	"fmt"

	"github.com/itouri/sgx-iaas/pkg/db/mongo"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"gopkg.in/mgo.v2/bson"
)

type ServiceInteractor struct {
	mongo.MongoHandler
	Collection string
}

func (i *ServiceInteractor) FindAll() (*[]keystone.Service, error) {
	res := new([]keystone.Service)
	err := i.MongoHandler.FindAll(i.Collection, res)
	if err != nil {
		return nil, fmt.Errorf("ServiceInteractor:FindAll() " + err.Error())
	}
	return res, nil
}

func (i *ServiceInteractor) FindByServiceType(serviceType keystone.EnumServiceType) (*keystone.Service, error) {
	query := bson.M{"type": serviceType}
	service := new(keystone.Service)
	// TODO FindOne?
	err := i.MongoHandler.FindOne(i.Collection, query, service)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (i *ServiceInteractor) FindByServiceID(serviceID string) (*keystone.Service, error) {
	query := bson.M{"service_id": serviceID}
	service := new(keystone.Service)
	err := i.MongoHandler.FindOne(i.Collection, query, service)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (i *ServiceInteractor) InsertService(service keystone.Service) error {
	return i.MongoHandler.Insert(i.Collection, service)
}

func (i *ServiceInteractor) DeleteService(serviceID string) error {
	query := bson.M{"service_id": serviceID}
	return i.MongoHandler.Delete(i.Collection, query)
}
