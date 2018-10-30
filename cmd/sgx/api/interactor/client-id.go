package interactor

import (
	"github.com/google/uuid"
	"github.com/itouri/sgx-iaas/pkg/db/mongo"
	"gopkg.in/mgo.v2/bson"
)

type ClientIDInteractor struct {
	*mongo.MongoHandler
	Collection string
}

var clientIDInteractor *ClientIDInteractor

func NewClientIDInteractor() *ClientIDInteractor {
	if clientIDInteractor == nil {
		clientIDInteractor = &ClientIDInteractor{
			MongoHandler: mongoHandler,
			Collection:   "client-id",
		}
	}
	return clientIDInteractor
}

// func (i *ClientIDInteractor) FindAll() (*[]uuid.UUID, error) {
// 	res := new([]uuid.UUID)
// 	err := i.MongoHandler.FindAll(i.Collection, "", res)
// 	if err != nil {
// 		return nil, fmt.Errorf("ClientIDInteractor:FindAll() " + err.Error())
// 	}
// 	return res, nil
// }

func (i *ClientIDInteractor) FindOneByCliendID(clientID uuid.UUID) (*uuid.UUID, error) {
	query := bson.M{"client_id": clientID}
	client := new(uuid.UUID)
	// TODO FindOne?
	err := i.MongoHandler.FindOne(i.Collection, query, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// func (i *ClientIDInteractor) FindByClientIDID(clientIDID string) (*uuid.UUID, error) {
// 	query := bson.M{"clientID_id": clientIDID}
// 	clientID := new(uuid.UUID)
// 	err := i.MongoHandler.FindOne(i.Collection, query, clientID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return clientID, nil
// }

func (i *ClientIDInteractor) InsertClientID(clientID uuid.UUID) error {
	return i.MongoHandler.Insert(i.Collection, clientID)
}

// func (i *ClientIDInteractor) UpsertClientID(clientID uuid.UUID) error {
// 	query := bson.M{"type": clientID.Type}
// 	upsert := bson.M{"$set": bson.M{"port": clientID.Port, "ipaddr": clientID.IPAddr}}
// 	return i.MongoHandler.Upsert(i.Collection, query, upsert)
// }

func (i *ClientIDInteractor) DeleteClientID(clientID string) error {
	query := bson.M{"clientID_id": clientID}
	return i.MongoHandler.Delete(i.Collection, query)
}
