package mq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeName = "notify"

	WechatKey   = "notify.wechat"
	DingTalkKey = "notify.dingtalk"

	UserIdKey = "userId"
)

func initConnection(amqpURI string) (conn *amqp.Connection, channel *amqp.Channel, err error) {
	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName("sample-consumer")
	fmt.Printf("dialing %q", amqpURI)
	conn, err = amqp.DialConfig(amqpURI, config)
	if err != nil {
		return nil, nil, fmt.Errorf("dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-conn.NotifyClose(make(chan *amqp.Error)))
	}()

	fmt.Printf("got Connection, getting Channel")
	channel, err = conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("channel: %s", err)
	}

	return conn, channel, nil
}

func declareDirect(channel *amqp.Channel, exchange, key string) error {
	err := channel.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		amqp.Table{},
	)
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare(key,
		true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = channel.QueueBind(key, key, exchange, false, nil)
	if err != nil {
		return err
	}

	return nil
}
