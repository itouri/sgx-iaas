package neutron

import (
	"fmt"
)

type EnumBridgeType int

const (
	BRIDGE_EXTERNAL EnumBridgeType = iota + 1
	BRIDGE_INTERNAL
)

func (b EnumBridgeType) String() (string, error) {
	name := ""
	switch b {
	case BRIDGE_EXTERNAL:
		name = "br-ex"
	case BRIDGE_INTERNAL:
		name = "br-int"
	default:
		return "", fmt.Errorf("unkown EnumBridgeType")
	}
	return name, nil
}
