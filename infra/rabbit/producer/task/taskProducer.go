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
	ExchangeName = "tasks"
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

	return &TaskProducer{conn, ch}
}

func (p *TaskProducer) CloseConnection() {
	p.connection.Close()
}

func (p *TaskProducer) CloseChannel() {
	p.channel.Close()
}

func (p *TaskProducer) DeclareExchange() error {
	err := p.channel.ExchangeDeclare(
		ExchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		log.Panicf("%s: %s", "Failed to declare an exchange", err)
	}
	return err
}

func (p *TaskProducer) Publish(mssge string) error {
	defer p.CloseConnection()
	defer p.CloseChannel()

	p.DeclareExchange()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.channel.PublishWithContext(ctx,
		ExchangeName, // exchange
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
