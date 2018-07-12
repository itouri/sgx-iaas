package amqp

type AMQPDriver struct {
	Url string
}

func NewAMQPDriver() *AMQPDriver {
	return &AMQPDriver{}
}

func (d *AMQPDriver) Send() {

}

func (d *AMQPDriver) SendNotification() {

}

func (d *AMQPDriver) Listen() {

}

func (d *AMQPDriver) ListenNotification() {

}

func (d *AMQPDriver) CleanUp() {

}

// getExchange
// getConnection
// getReplyQ Qって何?
