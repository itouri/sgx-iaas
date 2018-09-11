package agent

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/keystone/util"
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
			if err := alarmToHeat(); err != nil {
				return err
			}
		}
	}
	return nil
}

func alarmToHeat() error {
	// ???
	//msg := []byte(alarm.ID.String())
	//notifier.Send(msg)

	endpointURL := ""
	heatURL, err := util.ResolveServiceEndpoint(endpointURL, keystone.Heat)
	if err != nil {
		return err
	}

	resp, err := http.Post(heatURL, "application/json", nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	return nil
}
