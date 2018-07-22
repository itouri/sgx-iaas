package neutron

import "net"

type FloatingIP struct {
	IPAddr      net.IP
	FixedIPAddr net.IP
	Active      bool
}
