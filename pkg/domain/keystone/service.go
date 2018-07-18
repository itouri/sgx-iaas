package keystone

type EnumType int

const (
	Compute EnumType = iota + 1
	Ec2
	Identity
	Image
	Network
	Volume
)

type Service struct {
	Enabled   bool
	ID        string
	Name      string
	Type      EnumType //TODO to enum
	Links     string   // need not?
	EndPoints []EndPoint
}
