package main

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)


func ohpointninePost(cmd *cobra.Command, args []string) {

	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strings.Join(args, " ")),
		},
	)
	if err != nil {
		log.Fatal("Sending message:", err)
	}
}


func ohpointnineWatch(cmd *cobra.Command, args []string) {
	
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to receive message:", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Message received: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
