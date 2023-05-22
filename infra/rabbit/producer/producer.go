package producer

import (
	"fmt"
	"log"

	"github.com/ManuelP84/calendar/infra/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitProducer struct {
	connection *amqp.Connection
}

func NewRabbitProducer(settings *rabbit.RabbitSettings) *RabbitProducer {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s", settings.User, settings.Password, settings.Host, settings.Port)
	conn, err := amqp.Dial(addr)

	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	return &RabbitProducer{conn}
}
