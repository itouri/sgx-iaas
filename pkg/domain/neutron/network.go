package neutron

import uuid "github.com/satori/go.uuid"

type Network struct {
	ID   uuid.UUID
	Name string
}
