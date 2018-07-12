package notify

import (
	"github.com/streadway/amqp"
)

type RabbitNotifyClient struct {
	Url       string
	QueueName string
	channel   *amqp.Channel
}

func NewRabbitNotifyClient(url string, queueName string) *RabbitNotifyClient {
	return &RabbitNotifyClient{
		Url:       url,
		QueueName: queueName,
	}
}

func (r *RabbitNotifyClient) Start() error {
	conn, err := amqp.Dial(r.Url)
	if err != nil {
		defer conn.Close()
		return err
	}

	r.channel, err = conn.Channel()
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

func (r *RabbitNotifyClient) Send(body []byte) error {
	err := r.channel.Publish(
		r.QueueName, // exchange
		"",          // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	return err
}
