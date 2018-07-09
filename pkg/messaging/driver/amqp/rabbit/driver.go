package rabbit

import (
	"github.com/itouri/sgx-iaas/pkg/messaging/driver/amqp"
)

type RabbitDriver struct {
	AMQPDriver amqp.Driver
}
