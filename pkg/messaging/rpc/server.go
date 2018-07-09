package rpc

import (
	"github.com/itouri/sgx-iaas/pkg/messaging"
)

type RPCServer struct {
	MessageHanglingServer messaging.Server
	Target                Target
}

//tansport target endpoints
func NewRPCServer(tp RPCTransport, tg messaging.Target) *RPCServer {
	return &RPCServer{
		MessageHanglingServer: server,
		Target:                tg,
	}
}
