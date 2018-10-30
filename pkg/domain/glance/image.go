package glance

import "github.com/google/uuid"

// image構造体は必要 中にはimageの状態を格納
type Image struct {
	ID   uuid.UUID
	Data []byte //TODO convert to URL
}
