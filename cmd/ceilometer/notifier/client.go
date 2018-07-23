package notifier

import "github.com/itouri/sgx-iaas/pkg/messaging/rabbit/notify"

var client notify.RabbitNotifyClient

func init() {
	url := ""
	qName := "telemetry"
	client = *notify.NewRabbitNotifyClient(url, qName)
	client.Start()
}

func Send(bin []byte) {
	client.Send(bin)
}
