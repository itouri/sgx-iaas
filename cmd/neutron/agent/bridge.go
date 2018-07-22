package agent

import "os/exec"

const (
	BRIDGE_EXTERNAL_NAME = "br-ex"
	BRIDGE_INTERNAL_NAME = "br-int"
)

func init() {
	makeInitialBridge()
}

func makeInitialBridge() {
	exec.Command("ovs-vsctl", "add-br", BRIDGE_INTERNAL_NAME)
	exec.Command("ovs-vsctl", "add-br", BRIDGE_EXTERNAL_NAME)
}
