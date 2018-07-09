package rpc

import (
	"github.com/itouri/sgx-iaas/pkg/messaging"
)

type RPCTransport struct {
	Transport messaging.Transport
}

func NewRPCTransport() {

	transport := messaging.NewTransport() //driver
	return &RPCTransport{
		Transport: transport,
	}
}
