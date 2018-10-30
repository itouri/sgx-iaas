package neutron

import "github.com/google/uuid"

type Network struct {
	ID   uuid.UUID
	Name string
}
