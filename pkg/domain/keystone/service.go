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

func (s EnumServiceType) String() string {
	names := [...]string{
		"Compute",
		"Identity",
		"Image",
		"Network",
		"Volume",
	}
	if s < Compute || Volume < s {
		return ""
	}
	return names[s]
}

func ToEnumServiceType(str string) EnumServiceType {
	switch str {
	case "Compute":
		return Compute
	case "Identity":
		return Identity
	case "Image":
		return Identity
	case "Network":
		return Network
	case "Volume":
		return Volume
	default:
		return -1
	}
}

type Service struct {
	ID      uuid.UUID       `json:"id"`
	Enabled bool            `json:"enabled"`
	Name    string          `json:"name"`
	Type    EnumServiceType `json:"type"` //TODO to enum
	// Links     string   // need not?
	//EndPoints []EndPoint
	// URL    string URLにしたら誰がDNSするんだ？
	Port   uint   `json:"port"`
	IPAddr net.IP `json:"ip_address"`
}
