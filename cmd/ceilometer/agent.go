package ceilometer

import (
	"github.com/itouri/sgx-iaas/cmd/ceilometer/agent"
	"github.com/itouri/sgx-iaas/pkg/messaging/rabbit/notify"
)

func main() {
	url := "localhost"
	qName := "ceilometer"
	server := notify.NewRabbitNotifyServer(url, qName)
	server.Start()
	server.Listen(agent.Collector)
}
