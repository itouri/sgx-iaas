package interactor

import (
	"fmt"

	"github.com/itouri/sgx-iaas/pkg/db/mongo"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"gopkg.in/mgo.v2/bson"
)

type EndPointInteractor struct {
	mongo.MongoHandler
	Collection string
}

func (i *EndPointInteractor) FindAll() (*[]keystone.EndPoint, error) {
	res := new([]keystone.EndPoint)
	err := i.MongoHandler.FindAll(i.Collection, res)
	if err != nil {
		return nil, fmt.Errorf("EndPointInteractor:FindAll() " + err.Error())
	}
	return res, nil
}

func (i *EndPointInteractor) FindByEndPointID(endPointID string) (*keystone.EndPoint, error) {
	query := bson.M{"end_point_id": endPointID}
	endPoint := new(keystone.EndPoint)
	err := i.MongoHandler.FindOne(i.Collection, query, endPoint)
	if err != nil {
		return nil, err
	}
	return endPoint, nil
}

func (i *EndPointInteractor) InsertEndPoint(endPoint keystone.EndPoint) error {
	return i.MongoHandler.Insert(i.Collection, endPoint)
}

func (i *EndPointInteractor) DeleteEndPoint(endPointID string) error {
	query := bson.M{"end_point_id": endPointID}
	return i.MongoHandler.Delete(i.Collection, query)
}
