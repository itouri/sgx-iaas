package keystone

import (
	"bytes"
	"encoding/json"

	"github.com/google/uuid"
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
	RAKey
)

type Service struct {
	ID      uuid.UUID `json:"id"`
	Enabled bool      `json:"enabled"`
	// Name    string          `json:"name"`
	Type EnumServiceType `json:"type"` //TODO to enum
	// Links     string   // need not?
	//EndPoints []EndPoint
	// URL    string URLにしたら誰がDNSするんだ？
	Port   uint64 `json:"port"`
	IPAddr string `json:"ipaddr"`
}

var serviceID = map[EnumServiceType]string{
	Compute:    "Compute",
	Ceilometer: "Ceilometer",
	Glance:     "Glance",
	Newtron:    "Newtron",
	Nova:       "Nova",
	Heat:       "Heat",
	RA:         "RA",
	RAKey:      "RAKey",
}

var serviceName = map[string]EnumServiceType{
	"Compute":    Compute,
	"Ceilometer": Ceilometer,
	"Glance":     Glance,
	"Newtron":    Newtron,
	"Nova":       Nova,
	"Heat":       Heat,
	"RA":         RA,
	"RAKey":      RAKey,
}

// TODO もっといい方法がある
func (s EnumServiceType) String() string {
	return serviceID[s]
}

func ToEnumServiceType(str string) EnumServiceType {
	return serviceName[str]
}

func (s *EnumServiceType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(serviceID[*s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *EnumServiceType) UnmarshalJSON(b []byte) error {
	// unmarshal as string
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	// lookup value
	*s = serviceName[str]
	return nil
}
