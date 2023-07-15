package main

import (
	"fmt"

	"github.com/mdiaas/events-golang/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()
	messages := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, messages, "myqueue")
	for message := range messages {
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}
