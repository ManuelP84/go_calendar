package task

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ManuelP84/calendar/infra/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeName = "taskExchange"
	//queueName     = "taskQueue"
	//routingKey    = "taskEvents"
	//exchangeType  = "fanout"
	//prefetchCount = 10
)

type TaskProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewTaskProducer(settings *rabbit.RabbitSettings) *TaskProducer {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s", settings.User, settings.Password, settings.Host, settings.Port)
	conn, err := amqp.Dial(addr)

	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panicf("%s: %s", "Failed to declare an exchange", err)
	}

	return &TaskProducer{conn, ch}
}

func (p *TaskProducer) Publish(mssge string) error {
	//defer p.CloseConnection()
	//defer p.CloseChannel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.channel.PublishWithContext(ctx,
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(mssge),
		})

	if err != nil {
		log.Panicf("%s: %s", "Failed to publish a message", err)
	}
	return err
}
