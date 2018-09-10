package agent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/itouri/sgx-iaas/pkg/domain/ceilometer"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
)

// メッセージキューから来たデータを処理する
func Collector(b []byte) error {
	// []byteを構造体へキャスト
	telemetry := &ceilometer.Telemetry{}
	err := json.Unmarshal(b, telemetry)
	if err != nil {
		return err
	}

	// 登録されたStackのTemplateと比較
	return compare(telemetry)
}

func compare(tel *ceilometer.Telemetry) error {
	//TODO so dirty!!!
	for _, alarm := range registeredAlarms {
		var value float32
		alarming := false
		switch alarm.MeterName {
		case "cpu":
			value = tel.CPUUsage
		case "mem":
			value = tel.RAMUsage
		case "sgx_mem":
			value = tel.SGXRAMUsage
		}

		switch alarm.ComparisonOperator {
		case "Ge":
			if value >= alarm.Threshold {
				alarming = true
			}
		case "Le":
			if value <= alarm.Threshold {
				alarming = true
			}
		case "Gt":
			if value > alarm.Threshold {
				alarming = true
			}
		case "Lt":
			if value < alarm.Threshold {
				alarming = true
			}
		case "Eq":
			if value == alarm.Threshold {
				alarming = true
			}
		case "Ne":
			if value != alarm.Threshold {
				alarming = true
			}
		}

		if alarming {
			// ???
			//msg := []byte(alarm.ID.String())
			//notifier.Send(msg)

			endpointURL := ""
			// heatに情報を送るためにはendpointからIPを解決する必要がある
			resp, err := http.Get(endpointURL + "/services/resolve/" + "heat")
			if err != nil {
				return err
			}

			service := &keystone.Service{}
			err = decodeJSON(resp, service)
			if err != nil {
				return err
			}

			heatURL := "http://" + service.IPAddr.String() + ":" + string(service.Port) + "/"

			resp, err = http.Post(heatURL, "application/json", nil)
			if err != nil {
				return err
			}
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("status code is %d", resp.StatusCode)
			}
		}
	}
}

func decodeJSON(resp *http.Response, v interface{}) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
