package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queueName string) error {
	messages, err := ch.Consume(
		queueName,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for message := range messages {
		out <- message
	}
	return nil
}

func Publish(ch *amqp.Channel, body, exName string) error {
	err := ch.Publish(
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
