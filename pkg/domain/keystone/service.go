package keystone

import (
	"net"

	uuid "github.com/satori/go.uuid"
)

type EnumServiceType int

const (
	Compute EnumServiceType = iota + 1
	// Ec2
	Ceilometer
	Glance
	Newtron
	Nova
	Heat
	RA
)

func (s EnumServiceType) String() string {
	names := [...]string{
		"Compute",
		"Ceilometer",
		"Glance",
		"Newtron",
		"Nova",
		"Heat",
		"RA",
	}
	if s < Compute || RA < s {
		return ""
	}
	return names[s]
}

func ToEnumServiceType(str string) EnumServiceType {
	switch str {
	case "Compute":
		return Compute
	case "Ceilometer":
		return Ceilometer
	case "Glance":
		return Glance
	case "Newtron":
		return Newtron
	case "Nova":
		return Nova
	case "Heat":
		return Heat
	case "RA":
		return RA
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
