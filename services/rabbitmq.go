package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Auth-CRUD-REST/config"
	"github.com/streadway/amqp"
)

type Message struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

func PublishMessage(message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = config.RabbitMQChannel.Publish(
		"",
		config.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	if err != nil {
		return err
	}

	return nil
}

func ConsumeMessage() {
	msgs, err := config.RabbitMQChannel.Consume(
		config.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s \n", d.Body)
			var message Message
			err := json.Unmarshal(d.Body, &message)
			fmt.Println(message.Action)
			if err != nil {
				fmt.Println("Error decoding message: ", err)
				continue
			}
			switch message.Action {
			case "CreateUser":
				fmt.Println("User created")
			case "UpdateUser":
				fmt.Println("User updated")
			default:
				fmt.Println("Unknown action:", message.Action)
			}
		}
	}()

	fmt.Println("Waiting for messages...")
	<-make(chan bool)
}
