package agent

import (
	"os/exec"

	"github.com/itouri/sgx-iaas/pkg/domain/neutron"
	uuid "github.com/satori/go.uuid"
)

var ports []neutron.Port

func init() {
	ports = []neutron.Port{}
}

func MakePort(bridgeType neutron.PortEnumBridgeType, networkUUID uuid.UUID) *neutron.Port {
	// bridgeID と NetworkID をveth名として持った veth pairを作る
	nid := networkUUID.String()
	bridgeID := uuid.Must(uuid.NewV4())
	exec.Command("ip", "link", "add", bridgeID.String(), "type", "veth", "peer", "name", nid)

	port := &neutron.Port{
		ID:         uuid.Must(uuid.NewV4()),
		BridgeType: bridgeType,
		BridgeID:   bridgeID,
		NetworkID:  networkUUID,
	}
	addPort(port)
	return port
}

func ConnectPort(bridgeType neutron.EnumBridgeType, network neutron.Network) error {
	port := MakePort(bridgeType, network.ID)
	bid := port.BridgeID.String()
	nid := port.NetworkID.String()
	// veth pair を bridge に接続
	//TODO tag
	exec.Command("ovs-vsctl", "add-port", bid, nid)
	// veth pair を namespace に移す
	exec.Command("ip", "link", "set", nid, "netns", nid)

	//TODO ip address を付与
	// exec.Command("ip", "netns", "exec", nid, "ifconfig", nid, IPAddr)

	// veth pair を活性化
	exec.Command("ip", "link", "set", bid, "up")
	exec.Command("ip", "link", "set", nid, "up")

	//TODO needed?
	exec.Command("ip", "netns", "exec", nid, "ip", "link", "set", "lo", "up")
	return nil
}

func DisconnectPort(bridgeType neutron.EnumBridgeType, network neutron.Network) {

}

func addPort(port *neutron.Port) {
	ports = append(ports, *port)
}
