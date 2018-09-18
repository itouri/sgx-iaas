package compute

import (
	"encoding/json"

	"github.com/shirou/gopsutil/mem"

	"github.com/itouri/sgx-iaas/pkg/domain/ceilometer"
	"github.com/itouri/sgx-iaas/pkg/messaging/rabbit/notify"
)

var (
	meminfoPath string
	cpuinfoPath string
)

func init() {
	//TODO get from config file
	meminfoPath = ""
	cpuinfoPath = ""
}

func main() {
	// endpointからceilometerのアクセスポイントを解決
	url := ""
	qName := "ceilometer"
	client := notify.NewRabbitNotifyClient(url, qName)
	client.Start()

	publish(client)
}

func publish(client *notify.RabbitNotifyClient) error {
	// meminfoFile, err := os.Open(meminfoPath)
	// if err != nil {
	// 	return err
	// }
	// defer meminfoFile.Close()

	// cpuinfoFile, err := os.Open(cpuinfoPath)
	// if err != nil {
	// 	return err
	// }
	// defer cpuinfoFile.Close()

	for {
		//tlmt := metering(meminfoFile, cpuinfoFile)

		// TODO is need CPU usage?
		v, _ := mem.VirtualMemory()
		ramUsage := float32(v.Available/v.Total) * 100

		// TODO how to metering SGX ram usage?
		tlmt := &ceilometer.Telemetry{
			RAMUsage: ramUsage,
		}

		bin, err := json.Marshal(tlmt)
		if err != nil {
			return err
		}
		client.Send(bin)
	}
}

// func metering(meminfoReader io.Reader, cpuinfoReader io.Reader) (*ceilometer.Telemetry, error) {
// 	meminfo, err := parseMemInfo(meminfoReader)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cpuinfo, err := parseCpuInfo(cpuinfoReader)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ramUsage := meminfo[""]
// 	sgxRamUsage := meminfo[""]
// 	cpuUsage := cpuinfo[""]

// 	tlmt := &ceilometer.Telemetry{
// 		RAMUsage:
// 	}

// 	return tlmt, nil
// }

// // copy from https://github.com/prometheus/node_exporter/blob/ebdd5241234b367ebc221a0d942b1183c8df70ab/collector/meminfo_linux.go
// func parseMemInfo(r io.Reader) (map[string]float64, error) {
// 	var (
// 		memInfo = map[string]float64{}
// 		scanner = bufio.NewScanner(r)
// 		re      = regexp.MustCompile(`\((.*)\)`)
// 	)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		parts := strings.Fields(line)
// 		fv, err := strconv.ParseFloat(parts[1], 64)
// 		if err != nil {
// 			return nil, fmt.Errorf("invalid value in meminfo: %s", err)
// 		}
// 		key := parts[0][:len(parts[0])-1] // remove trailing : from key
// 		// Active(anon) -> Active_anon
// 		key = re.ReplaceAllString(key, "_${1}")
// 		switch len(parts) {
// 		case 2: // no unit
// 		case 3: // has unit, we presume kB
// 			fv *= 1024
// 			key = key + "_bytes"
// 		default:
// 			return nil, fmt.Errorf("invalid line in meminfo: %s", line)
// 		}
// 		memInfo[key] = fv
// 	}
// 	return memInfo, scanner.Err()
// }
