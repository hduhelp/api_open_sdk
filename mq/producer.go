package mq

import (
	"context"
	"crypto/md5"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type Producer struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel

	appId string
}

func NewProducer(appId string, amqpURI string) (*Producer, error) {
	if appId == "" {
		return nil, fmt.Errorf("app id is empty")
	}

	p := &Producer{
		appId: appId,
	}

	var err error

	p.Conn, p.Channel, err = initConnection(amqpURI)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Producer) PublishNotice(ctx context.Context, data *NoticeTemplate, options ...string) error {

	if data == nil {
		return fmt.Errorf("data is empty")
	}

	opts := parseOptions(options)

	if data.WeChat != nil {
		for _, staffId := range data.StaffID {
			if staffId == "" {
				continue
			}
			body, err := MarshalWechatBody(staffId, data.WeChat)
			if err != nil {
				return err
			}
			if err := p.publish(ctx, WechatKey, body, opts); err != nil {
				return err
			}
		}
	}

	if data.DingTalk != nil {
		for _, staffId := range data.StaffID {
			if staffId == "" {
				continue
			}
			body, err := MarshalWechatBody(staffId, data.WeChat)
			if err != nil {
				return err
			}
			if err := p.publish(ctx, DingTalkKey, body, opts); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseOptions(options []string) map[string]string {
	opts := make(map[string]string)
	for i := 0; i < len(options); i += 2 {
		if i+1 < len(options) {
			opts[options[i]] = options[i+1]
		}
	}
	return opts
}

func (p *Producer) publish(ctx context.Context, key string, msg []byte, opts map[string]string) error {

	carrier := RabbitMQHeaderCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	headers := amqp.Table(carrier)

	err := p.Channel.PublishWithContext(
		ctx,
		exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         msg,
			AppId:        p.appId,
			UserId:       opts[UserIdKey],
			MessageId:    fmt.Sprintf("%x", md5.Sum(msg)),
			Headers:      headers,
		})

	return err

}

func (p *Producer) Shutdown() error {

	if err := p.Conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer fmt.Printf("AMQP shutdown OK")

	return nil
}
