package neutron

import "github.com/google/uuid"

type Router struct {
	// FloatingIPs []string
	ID   uuid.UUID
	Name string
}
