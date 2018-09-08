package glance

import uuid "github.com/satori/go.uuid"

// この構造体は必要ないと思う
type Image struct {
	ID   uuid.UUID
	Data []byte //TODO convert to URL
}
