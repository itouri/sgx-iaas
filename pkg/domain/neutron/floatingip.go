package neutron

import (
	"net"
)

type EnumStatus int

const (
	Active EnumStatus = iota + 1
	Down
)

type FloatingIP struct {
	Address        net.IP
	FixedIPAddress net.IP
	Status         EnumStatus
	// RouterID       string
}
