package messaging

import (
	"github.com/itouri/sgx-iaas/pkg/messaging/driver"
)

type Transport struct {
	Driver driver.Driver
	// Conf driver.conf
}

// conf
func NewTansport() *Transport {

	return &Transport{
		Driver: driver,
	}
}

func (t *Transport) send() {

}

func (t *Transport) sendNotification() {

}

func (t *Transport) listen() {

}

func (t *Transport) listenNotification() {

}

func (t *Transport) CleanUp() {

}
