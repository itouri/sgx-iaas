package compute

import (
	"encoding/json"
	"fmt"

	"github.com/itouri/sgx-iaas/pkg/domain/ceilometer"
	"github.com/itouri/sgx-iaas/pkg/messaging/rabbit/notify"
)

func main() {
	// endpointからceilometerのアクセスポイントを解決
	url := ""
	qName := "ceilometer"
	client := notify.NewRabbitNotifyClient(url, qName)
	client.Start()

}

func publisher(client *notify.RabbitNotifyClient) {
	for {
		tlmt := ceilometer.Telemetry{}
		bin, err := json.Marshal(tlmt)
		if err != nil {
			fmt.Println(err)
			return
		}
		client.Send(bin)
	}
}
