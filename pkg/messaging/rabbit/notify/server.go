package notify

import (
	"github.com/streadway/amqp"
)

type ListenFunc func([]byte)

type RabbitNotifyServer struct {
	Url       string
	QueueName string
	channel   *amqp.Channel
}

func NewRabbitNotifyServer(url string, queueName string) (*RabbitNotifyServer, error) {
	return nil, &RabbitNotifyServer{
		Url:       url,
		QueueName: queueName,
		channel:   ch,
	}
}

func (*r RabbitNotifyServer) Start() error 
{
	conn, err := amqp.Dial(url)
	if err != nil {
		defer conn.Close()
		return err
	}

	r.ch, err := conn.Channel()
	if err != nil {
		defer ch.Close()
		return err
	}

	err = r.ch.ExchangeDeclare(
		r.QueueName,   // name
		"topic", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func (*r RabbitNotifyServer) Listen(lf ListenFunc) error 
{
	q, err := r.ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = r.ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		r.QueueName, // exchange
		false,
		nil
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	go func() {
		for msg := range msgs {
			lf(msg)
		}
	}
}