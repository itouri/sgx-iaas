package notify

import (
	"github.com/streadway/amqp"
)

type RabbitNotifyClient struct {
	Url       string
	QueueName string
	channel   *amqp.Channel
}

func NewRabbitNotifyClient(url string, queueName string) (*RabbitNotifyClient, error) {
	return nil, &RabbitNotifyClient{
		Url:       url,
		QueueName: queueName,
		channel:   ch,
	}
}

func (*r RabbitNotifyClient) Start() error 
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

func (*r RabbitNotifyClient) Send(body []byte) error 
{
	err = ch.Publish(
		r.QueueName, // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
		}
	)
}