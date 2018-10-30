package agent

import (
	"os/exec"

	"github.com/google/uuid"
	"github.com/itouri/sgx-iaas/pkg/domain/neutron"
)

var routers []neutron.Router
var externalRouter neutron.Router

func init() {
	routers = []neutron.Router{}
}

// TODO execでやるしかない？
func MakeRouter(name string) uuid.UUID {
	// add namespace
	id := uuid.Must(uuid.NewV4())
	idstr := id.String()
	exec.Command("ip", "netns", "add", idstr)

	router := &neutron.Router{
		ID:   id,
		Name: name,
	}
	addRouter(router)
	return id
}

func DeleteRouter(id uuid.UUID) {
	delIndex := -1
	for i, r := range routers {
		if r.ID == id {
			delIndex = i
			break
		}
	}
	if delIndex == -1 {
		return
	}
	routers = append(routers[:delIndex], routers[delIndex+1:]...)
	exec.Command("ip", "netns", "delete", id.String())
}

func GetRouter() []neutron.Router {
	return routers
}

func addRouter(router *neutron.Router) {
	routers = append(routers, *router)
}
