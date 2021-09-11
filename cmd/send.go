package main

import (
	"github.com/streadway/amqp"
	"io/ioutil"
	"test/logging"
)

func main() {
	logger := logging.GetLogger()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	{
		if err != nil {
			logger.Error(err)
		}
	}
	defer conn.Close()
	logger.Info("connected to rabbit")

	ch, err := conn.Channel()
	if err != nil {
		logger.Error("failed to open a channel")
	}
	defer ch.Close()

	logger.Info("created queue")
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		logger.Error("failed to open connect to queue")
	}
	logger.Info("prepare to publish")

	body, err := ioutil.ReadFile("logs/all.log")
	err = ch.Publish(
		"events", // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		logger.Error("failed to publish a message")
	}

}
