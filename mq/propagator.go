package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQHeaderCarrier adapts http.Header to satisfy the TextMapCarrier interface.
type RabbitMQHeaderCarrier amqp.Table

// Get returns the value associated with the passed key.
func (hc RabbitMQHeaderCarrier) Get(key string) string {
	return hc.Get(key)
}

// Set stores the key-value pair.
func (hc RabbitMQHeaderCarrier) Set(key string, value string) {
	//hc.Set(key, value)
	hc[key] = value
}

// Keys lists the keys stored in this carrier.
func (hc RabbitMQHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range hc {
		keys = append(keys, k)
	}
	return keys
}
