package agent

import (
	"github.com/itouri/sgx-iaas/pkg/domain/heat"
	uuid "github.com/satori/go.uuid"
)

var registeredAlarms []heat.Alarm

func init() {
	registeredAlarms = []heat.Alarm{}
}

func RegisterAlarm(alarm *heat.Alarm) {
	registeredAlarms = append(registeredAlarms, *alarm)
}

func DeleteAlarm(id uuid.UUID) {
	delIndex := -1
	for i, r := range registeredAlarms {
		if r.ID == id {
			delIndex = i
			break
		}
	}
	if delIndex == -1 {
		return
	}
	registeredAlarms = append(registeredAlarms[:delIndex], registeredAlarms[delIndex+1:]...)
}

func GetRegisteredAlarms() []heat.Alarm {
	return registeredAlarms
}
