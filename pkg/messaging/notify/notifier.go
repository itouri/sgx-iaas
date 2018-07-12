package notify

import (
	"github.com/itouri/sgx-iaas/pkg/messaging"
)

type Notifier struct {
	Transport messaging.Transport
	Topic     string
}

func NewNotifier(tp messaging.Transport, topic string) *Notifier {
	return &Notifier{
		Transport: tp,
		Topic:     topic,
	}
}
