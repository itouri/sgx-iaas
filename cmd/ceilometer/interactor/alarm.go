package interactor

import (
	"fmt"

	"github.com/itouri/sgx-iaas/pkg/db/mongo"
	"github.com/itouri/sgx-iaas/pkg/domain/heat"
	"gopkg.in/mgo.v2/bson"
)

type AlarmInteractor struct {
	mongo.MongoHandler
	Collection string
}

func (i *AlarmInteractor) FindAll() (*[]heat.Alarm, error) {
	res := new([]heat.Alarm)
	err := i.MongoHandler.FindAll(i.Collection, res)
	if err != nil {
		return nil, fmt.Errorf("AlarmInteractor:FindAll() " + err.Error())
	}
	return res, nil
}

func (i *AlarmInteractor) FindByAlarmID(alarmID string) (*heat.Alarm, error) {
	query := bson.M{"alarm_id": alarmID}
	alarm := new(heat.Alarm)
	err := i.MongoHandler.FindOne(i.Collection, query, alarm)
	if err != nil {
		return nil, err
	}
	return alarm, nil
}

func (i *AlarmInteractor) InsertAlarm(alarm heat.Alarm) error {
	return i.MongoHandler.Insert(i.Collection, alarm)
}

func (i *AlarmInteractor) DeleteAlarm(alarmID string) error {
	query := bson.M{"alarm_id": alarmID}
	return i.MongoHandler.Delete(i.Collection, query)
}
