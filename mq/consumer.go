package mq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
	Tag        string
	Done       chan error
	Deliveries <-chan amqp.Delivery
}

func NewConsumer(amqpURI, key string) (*Consumer, error) {
	c := &Consumer{
		Conn:    nil,
		Channel: nil,
		Tag:     key,
		Done:    make(chan error),
	}

	var err error

	c.Conn, c.Channel, err = initConnection(amqpURI)
	if err != nil {
		return nil, err
	}

	err = declareDirect(c.Channel, exchangeName, key)
	if err != nil {
		return nil, err
	}

	c.Deliveries, err = c.Channel.Consume(
		key,   // name
		key,   // consumerTag,
		false, // autoAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("queue Consume: %s", err)
	}

	return c, nil
}

func (c *Consumer) Shutdown() error {
	// will close() the Deliveries Channel
	if err := c.Channel.Cancel(c.Tag, true); err != nil {
		return fmt.Errorf("consumer cancel failed: %s", err)
	}

	if err := c.Conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer fmt.Printf("AMQP shutdown OK")

	// wait for exit
	return <-c.Done
}
