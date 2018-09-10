package ceilometer

// alarmの情報をDBに保存する必要がある
// k8sもKVSだよなたしか

// computeからtelemetryへ送る情報
type Telemetry struct {
	CPUUsage    float32 `json:"cpu-usage"`
	RAMUsage    float32 `json:"ram-usage"`
	SGXRAMUsage float32 `json:"sgx-ram-usage"`
}
