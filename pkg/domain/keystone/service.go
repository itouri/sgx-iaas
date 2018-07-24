package keystone

import (
	"net"

	uuid "github.com/satori/go.uuid"
)

type EnumServiceType int

const (
	Compute EnumServiceType = iota + 1
	// Ec2
	Identity
	Image
	Network
	Volume
)

type Service struct {
	ID      uuid.UUID
	Enabled bool
	Name    string
	Type    EnumServiceType //TODO to enum
	// Links     string   // need not?
	//EndPoints []EndPoint
	// URL    string URLにしたら誰がDNSするんだ？
	IPAddr net.IP
}
