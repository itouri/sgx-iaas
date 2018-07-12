package rabbit

import (
	"github.com/itouri/sgx-iaas/pkg/messaging/driver/amqp"
)

type RabbitDriver struct {
	AMQPDriver *amqp.AMQPDriver
}

func NewRabbitDriver() *RabbitDriver {
	amqpDriver := amqp.NewAMQPDriver()
	return &RabbitDriver{
		AMQPDriver: amqpDriver,
	}
}
