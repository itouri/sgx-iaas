package interactor

import (
	"fmt"

	"github.com/itouri/sgx-iaas/pkg/db/mongo"
	"github.com/itouri/sgx-iaas/pkg/domain/heat"
	"gopkg.in/mgo.v2/bson"
)

type TemplateInteractor struct {
	mongo.MongoHandler
	Collection string
}

func (i *TemplateInteractor) FindAll() (*[]heat.Template, error) {
	res := new([]heat.Template)
	err := i.MongoHandler.FindAll(i.Collection, res)
	if err != nil {
		return nil, fmt.Errorf("TemplateInteractor:FindAll() " + err.Error())
	}
	return res, nil
}

func (i *TemplateInteractor) FindByTemplateID(TemplateID string) (*heat.Template, error) {
	query := bson.M{"Template_id": TemplateID}
	Template := new(heat.Template)
	err := i.MongoHandler.FindOne(i.Collection, query, Template)
	if err != nil {
		return nil, err
	}
	return Template, nil
}

func (i *TemplateInteractor) InsertTemplate(Template heat.Template) error {
	return i.MongoHandler.Insert(i.Collection, Template)
}

func (i *TemplateInteractor) DeleteTemplate(TemplateID string) error {
	query := bson.M{"Template_id": TemplateID}
	return i.MongoHandler.Delete(i.Collection, query)
}
