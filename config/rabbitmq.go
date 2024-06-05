package config

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	RabbitMQConn    *amqp.Connection
	RabbitMQChannel *amqp.Channel
	Queue           amqp.Queue
)

func InitRabbitMQ() {
	var err error
	rabbitMQUser := os.Getenv("RABBITMQ_USERNAME")
	rabbitMQPass := os.Getenv("RABBITMQ_PASSWORD")
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQUser, rabbitMQPass, rabbitMQHost, rabbitMQPort)
	RabbitMQConn, err = amqp.Dial(connectionString)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	Queue, err = RabbitMQChannel.QueueDeclare(
		"user_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
