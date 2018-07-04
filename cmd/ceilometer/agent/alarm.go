package agent

import (
	"github.com/itouri/sgx-iaas/pkg/domain/heat"
)

var registeredAlarms []heat.Alarm

func init() {
	registeredAlarms = []heat.Alarm{}
}

func RegisterAlarm(alarm *heat.Alarm) {
	registeredAlarms = append(registeredAlarms, *alarm)
}

func GetRegisteredAlarms() []heat.Alarm {
	return registeredAlarms
}
