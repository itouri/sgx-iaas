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
	ID      uuid.UUID       `json:"id"`
	Enabled bool            `json:"enabled"`
	Name    string          `json:"name"`
	Type    EnumServiceType `json:"type"` //TODO to enum
	// Links     string   // need not?
	//EndPoints []EndPoint
	// URL    string URLにしたら誰がDNSするんだ？
	IPAddr net.IP `json:"ip_address"`
}
