package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := 	consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (c * Consumer) setup() error {
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}
	
	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
