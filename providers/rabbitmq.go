package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

type HandleRabbitMQFunc func(message []byte) (bool, error)

type RabbitMQConfig struct {
	Address  string
	Port     int
	Username string
	Password string
	Topics   []string
}

type IRabbitMQ interface {
	Publish(ctx context.Context, topic string, message interface{}) error
	Consume(topic string, handleHandleRabbitMQFunc HandleRabbitMQFunc) error
}

type rabbitMQ struct {
	channel *amqp091.Channel
}

func NewRabbitMQ(config RabbitMQConfig) (IRabbitMQ, func()) {
	conn, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Address, config.Port))
	if err != nil {
		panic(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	for _, topic := range config.Topics {
		_, err = channel.QueueDeclare(topic, true, false, false, false, nil)
		if err != nil {
			panic(err)
		}
	}
	return &rabbitMQ{
			channel: channel,
		}, func() {
			_ = channel.Close()
			_ = conn.Close()
		}
}

func (rabbit *rabbitMQ) Consume(topic string, handleHandleRabbitMQFunc HandleRabbitMQFunc) error {
	messages, err := rabbit.channel.Consume(topic, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	for d := range messages {
		if ok, err := handleHandleRabbitMQFunc(d.Body); err != nil {
			return err
		} else if !ok {
		}
	}
	return nil
}

func (rabbit *rabbitMQ) Publish(ctx context.Context, topic string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return rabbit.channel.PublishWithContext(ctx, "", topic, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
