package keystone

import (
	"github.com/satori/go.uuid"
)

// type EnumInterface int

// const (
// 	Admin EnumInterface = iota + 1
// 	Internal
// 	Public
// )

type EndPoint struct {
	ID      uuid.UUID // TODO UUID?
	Enabled bool
	// Interface EnumInterface // TODO to enum
	ServiceID string
	URL       string
}
