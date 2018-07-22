package neutron

import (
	"net"

	"github.com/satori/go.uuid"
)

type Port struct {
	ID uuid.UUID
	// Name      string
	BridgeType     EnumBridgeType
	BridgeID       uuid.UUID
	NetworkID      uuid.UUID
	NetworkIPAdder net.IP // network側portに紐付いているip address
}
