package agent

import (
	"net"
)

var ipAddresses []net.IP

func init() {
	ipAddresses = []net.IP{}
}
