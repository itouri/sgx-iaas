package amqp

type Driver struct {
}

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) Send() {

}

func (d *Driver) SendNotification() {

}

func (d *Driver) Listen() {

}

func (d *Driver) ListenNotification() {

}

func (d *Driver) CleanUp() {

}

// getExchange
// getConnection
// getReplyQ Qって何?
