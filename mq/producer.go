package mq

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/singleflight"
)

type Producer struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel

	amqpURI string // AMQP URI for RabbitMQ reconnection
	sf      *singleflight.Group

	appId string
}

func NewProducer(appId string, amqpURI string) (*Producer, error) {
	if appId == "" {
		return nil, fmt.Errorf("app id is empty")
	}

	p := &Producer{
		appId:   appId,
		amqpURI: amqpURI,
		sf:      new(singleflight.Group),
	}

	var err error

	p.Conn, p.Channel, err = initConnection(amqpURI)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Producer) isConnected() bool {
	return !p.Conn.IsClosed() && !p.Channel.IsClosed()
}

func (p *Producer) connectFn() error {
	if p.isConnected() {
		return nil
	}

	_, err, _ := p.sf.Do("reconnect", func() (interface{}, error) {
		var lastErr error
		for i := 0; i < 3; i++ {
			if p.isConnected() {
				return nil, nil
			}

			if i > 0 {
				time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i-1))))
			}

			conn, channel, err := initConnection(p.amqpURI)
			if err != nil {
				lastErr = fmt.Errorf("reconnect attempt %d failed: %s", i+1, err)
				continue
			}

			oldConn := p.Conn
			oldChannel := p.Channel
			p.Conn = conn
			p.Channel = channel

			if oldChannel != nil {
				_ = oldChannel.Close()
			}
			if oldConn != nil {
				_ = oldConn.Close()
			}

			return nil, nil
		}

		return nil, lastErr
	})

	return err
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
			body, err := MarshalDingTalkBody(staffId, data.DingTalk)
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

	if err != nil && errors.Is(err, amqp.ErrClosed) {
		if err = p.connectFn(); err != nil {
			return err
		}

		err = p.Channel.PublishWithContext(
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
	}

	return err

}

func (p *Producer) Shutdown() error {

	if err := p.Conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer fmt.Printf("AMQP shutdown OK")

	return nil
}
