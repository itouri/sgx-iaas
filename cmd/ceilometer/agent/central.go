package agent

import (
	"bytes"
	"encoding/binary"

	"github.com/itouri/sgx-iaas/cmd/ceilometer/notifier"
	"github.com/itouri/sgx-iaas/pkg/domain/ceilometer"
)

// メッセージキューから来たデータを処理する
func Collector(b []byte) {
	// []byteを構造体へキャスト
	telemetry := &ceilometer.Telemetry{}
	reader := bytes.NewReader(b)
	binary.Read(reader, binary.LittleEndian, telemetry)

	// 登録されたStackのTemplateと比較
	compare(telemetry)
}

func compare(tel *ceilometer.Telemetry) {
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
			msg := []byte(alarm.AlarmAction)
			notifier.Send(msg)

			//TODO
			// heatに情報を送るためにはendpointからIPを解決する必要がある
			// ip, port := http.Get(endpointURL + heat)
			// resp, err := http.Post(heatURL + /alarm.AlarmAction)
		}
	}
}
