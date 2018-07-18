package keystone

type EnumInterface int

const (
	Admin EnumInterface = iota + 1
	Internal
	Public
)

type EndPoint struct {
	ID        string // TODO UUID?
	Enabled   bool
	Interface EnumInterface // TODO to enum
	ServiceID string
	URL       string
}
