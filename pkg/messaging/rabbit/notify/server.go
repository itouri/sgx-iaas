package notify

import (
	"github.com/streadway/amqp"
)

type ListenFunc func([]byte)

type RabbitNotifyServer struct {
	Url       string
	QueueName string
	conn      *amqp.Connection
	channel   *amqp.Channel
}

func NewRabbitNotifyServer(url string, queueName string) *RabbitNotifyServer {
	return &RabbitNotifyServer{
		Url:       url,
		QueueName: queueName,
	}
}

//TODO connection, channelの管理方法
func (r *RabbitNotifyServer) Start() error {
	var err error
	r.conn, err = amqp.Dial(r.Url)
	if err != nil {
		defer r.conn.Close()
		return err
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		defer r.channel.Close()
		return err
	}

	err = r.channel.ExchangeDeclare(
		r.QueueName, // name
		"topic",     // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitNotifyServer) Stop() {
	r.conn.Close()
	r.channel.Close()
}

func (r *RabbitNotifyServer) Listen(lf ListenFunc) error {
	q, err := r.channel.QueueDeclare(
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

	err = r.channel.QueueBind(
		q.Name,      // queue name
		"",          // routing key
		r.QueueName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := r.channel.Consume(
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
			lf(msg.Body)
		}
	}()

	return nil
}
