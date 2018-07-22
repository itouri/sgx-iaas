package agent

import (
	"net"
)

// 割り当てたIPアドレスを管理する
var ipAddresses []net.IP

func init() {
	ipAddresses = []net.IP{}
}
