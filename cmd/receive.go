package main

import (
	"github.com/streadway/amqp"
	"log"
	"test/logging"
)

func main() {
	logger := logging.GetLogger()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	logger.Info("connected to rabbit")
	if err != nil {
		logger.Error("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Error("Failed to connect to channel")
	}
	defer ch.Close()

	logger.Info("connected to channel")

	ch.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	err = ch.QueueBind("hello", "#", "events", false, nil)

	logger.Info("binded queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)
	logger.Info("registered a consumer")
	go func() {
		for d := range msgs {
			log.Printf("\nReceived a message:\n%s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
