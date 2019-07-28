package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"pack.ag/amqp"
)


func onepointohPost(cmd *cobra.Command, args []string) {

	client, err := amqp.Dial(connectionString, amqp.ConnSASLPlain(username, password))
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()

	{
		sender, err := session.NewSender(amqp.LinkTargetAddress("/" + queueName))
		if err != nil {
			log.Fatal("Creating sender link:", err)
		}

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		err = sender.Send(ctx, amqp.NewMessage([]byte(strings.Join(args, " "))))
		if err != nil {
			log.Fatal("Sending message:", err)
		}

		sender.Close(ctx)
		cancel()
	}
}


func onepointohWatch(cmd *cobra.Command, args []string) {

	client, err := amqp.Dial(connectionString, amqp.ConnSASLPlain(username, password))
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()

	{
		receiver, err := session.NewReceiver(amqp.LinkSourceAddress("/"+queueName), amqp.LinkCredit(10))
		if err != nil {
			log.Fatal("Creating receiver link:", err)
		}

		defer func() {
			ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
			receiver.Close(ctx)
			cancel()
		}()

		for {
			msg, err := receiver.Receive(ctx)
			if err != nil {
				log.Fatal("Reading message from AMQP:", err)
			}

			msg.Accept()

			fmt.Printf("Message received: %s\n", msg.GetData())
		}

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	}
}
