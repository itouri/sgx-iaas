package neutron

import uuid "github.com/satori/go.uuid"

type Router struct {
	// FloatingIPs []string
	ID   uuid.UUID
	Name string
}
