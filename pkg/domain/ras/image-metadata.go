package ras

import (
	"github.com/satori/go.uuid"
)

type ImageMetadata struct {
	ClientID  uuid.UUID `bson:"client_id" json:"client_id"`
	ImageID   uuid.UUID `bson:"image_id" json:"image_id"`
	Nonce     string    `bson:"nonce" json:"nonce"`
	Available bool      `bson:"available" json:"available"`
}
