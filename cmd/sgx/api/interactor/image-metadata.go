package interactor

import (
	"fmt"

	"github.com/itouri/sgx-iaas/pkg/db/mongo"
	"github.com/itouri/sgx-iaas/pkg/domain/ras"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type ImageMetadataInteractor struct {
	*mongo.MongoHandler
	Collection string
}

var imageMetadataInteractor *ImageMetadataInteractor

func NewImageMetadataInteractor() *ImageMetadataInteractor {
	if imageMetadataInteractor == nil {
		imageMetadataInteractor = &ImageMetadataInteractor{
			MongoHandler: mongoHandler,
			Collection:   "image-metadata",
		}
	}
	return imageMetadataInteractor
}

func (i *ImageMetadataInteractor) FindAll() (*[]ras.ImageMetadata, error) {
	res := new([]ras.ImageMetadata)
	err := i.MongoHandler.FindAll(i.Collection, "", res)
	if err != nil {
		return nil, fmt.Errorf("ImageMetadataInteractor:FindAll() " + err.Error())
	}
	return res, nil
}

func (i *ImageMetadataInteractor) FindByClientID(clientID uuid.UUID) (*ras.ImageMetadata, error) {
	query := bson.M{"client_id": clientID}
	imageMetadata := new(ras.ImageMetadata)
	// TODO FindOne?
	err := i.MongoHandler.FindOne(i.Collection, query, imageMetadata)
	if err != nil {
		return nil, err
	}
	return imageMetadata, nil
}

// func (i *ImageMetadataInteractor) FindByImageMetadataID(imageMetadataID string) (*ras.ImageMetadata, error) {
// 	query := bson.M{"imageMetadata_id": imageMetadataID}
// 	imageMetadata := new(ras.ImageMetadata)
// 	err := i.MongoHandler.FindOne(i.Collection, query, imageMetadata)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return imageMetadata, nil
// }

func (i *ImageMetadataInteractor) InsertImageMetadata(imageMetadata *ras.ImageMetadata) error {
	return i.MongoHandler.Insert(i.Collection, imageMetadata)
}

// func (i *ImageMetadataInteractor) UpsertImageMetadata(imageMetadata ras.ImageMetadata) error {
// 	query := bson.M{"type": imageMetadata.Type}
// 	upsert := bson.M{"$set": bson.M{"port": imageMetadata.Port, "ipaddr": imageMetadata.IPAddr}}
// 	return i.MongoHandler.Upsert(i.Collection, query, upsert)
// }

func (i *ImageMetadataInteractor) DeleteImageMetadata(imageMetadataID string) error {
	query := bson.M{"imageMetadata_id": imageMetadataID}
	return i.MongoHandler.Delete(i.Collection, query)
}
