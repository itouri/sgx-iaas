package interactor

import (
	"github.com/itouri/sgx-iaas/pkg/db/mongo"
	"github.com/itouri/sgx-iaas/pkg/domain/heat"
	"gopkg.in/mgo.v2/bson"
)

type AlarmInteractor struct {
	mongo.MongoHandler
	Collection string
}

func (i *AlarmInteractor) FindByAlarmID(alarmID string) (*heat.Alarm, error) {
	// TODO テスト
	query := bson.M{"alarms": bson.M{"$elemMatch": bson.M{"alarm_id": alarmID}}}
	alarms := new([]heat.Alarm)
	err := i.MongoHandler.FindOne(i.Collection, query, alarms)
	if err != nil {
		return nil, err
	}

	var alarm *heat.Alarm
	for _, a := range *alarms {
		if a.ID.String() == alarmID {
			alarm = &a
		}
	}

	return alarm, nil
}
