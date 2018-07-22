package agent

import (
	"os/exec"

	"github.com/itouri/sgx-iaas/pkg/domain/neutron"
	uuid "github.com/satori/go.uuid"
)

var networks []neutron.Network

func init() {
	networks = []neutron.Network{}
}

func MakeNetwork(name string) uuid.UUID {
	// add namespace
	id := uuid.Must(uuid.NewV4())
	idstr := id.String()
	exec.Command("ip", "netns", "add", idstr)

	network := &neutron.Network{
		ID:   id,
		Name: name,
	}
	addNetwork(network)
	return id
}

func DeleteNetwork(id uuid.UUID) {
	delIndex := -1
	for i, r := range networks {
		if r.ID == id {
			delIndex = i
			break
		}
	}
	if delIndex == -1 {
		return
	}
	networks = append(networks[:delIndex], networks[delIndex+1:]...)
	exec.Command("ip", "netns", "delete", id.String())
}

func GetNetwork() []neutron.Network {
	return networks
}

func addNetwork(network *neutron.Network) {
	networks = append(networks, *network)
}
