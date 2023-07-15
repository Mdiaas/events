package main

import "github.com/mdiaas/events-golang/pkg/rabbitmq"

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()
	rabbitmq.Publish(ch, "Hello world!", "amq.direct")
}
