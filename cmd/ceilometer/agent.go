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
	// TODO error handle
	server.Listen(agent.Collector)
}
