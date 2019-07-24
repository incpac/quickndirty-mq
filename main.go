package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	
	"github.com/spf13/cobra"
	"pack.ag/amqp"
)


var Version string

var connectionString string
var username string
var password string
var queueName string


func post(cmd *cobra.Command, args []string) {

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
		sender, err := session.NewSender(amqp.LinkTargetAddress("/"+queueName))
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


func watch(cmd *cobra.Command, args []string) {

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
	}
}


func main() {

	command := &cobra.Command{
		Use:	"qndmq",
		Short:	"A quick 'n' dirty Apache Active MQ client",
		Long:	"A simple Apache Active MQ client useful for testing configurations and running servers",
		Run: func(cmd *cobra.Command, args[]string) {
			fmt.Printf("%s\n", Version)
		},
	}


	postCommand := &cobra.Command{
		Use:	"post [flags] [message]",
		Short:	"Post a message to the queue",
		Long: 	"Post a message to the queue",
		Args:	cobra.MinimumNArgs(1),
		Run: 	post,
	}

	postCommand.Flags().StringVarP(&connectionString,	"connection",	"c", 	os.Getenv("QNDMQ_CONNECTION"),	"The connection string for the Active MQ server")
	postCommand.Flags().StringVarP(&username, 		"username", 	"u", 	os.Getenv("QNDMQ_USERNAME"), 	"The username to connect to Active MQ with")
	postCommand.Flags().StringVarP(&password, 		"password", 	"p", 	os.Getenv("QNDMQ_PASSWORD"),	"The password to connect to Active MQ with")
	postCommand.Flags().StringVarP(&queueName, 		"queue", 	"q", 	os.Getenv("QNDMQ_QUEUE"), 	"The name of the queue to post to")

	command.AddCommand(postCommand)


	watchCommand := &cobra.Command{
		Use:	"watch",
		Short:	"Watch the queue for new messages",
		Long:	"Watch the queue for new messages",
		Run: 	watch,
	}

	watchCommand.Flags().StringVarP(&connectionString,	"connection",	"c", 	os.Getenv("QNDMQ_CONNECTION"),	"The connection string for the Active MQ server")
	watchCommand.Flags().StringVarP(&username, 		"username", 	"u", 	os.Getenv("QNDMQ_USERNAME"), 	"The username to connect to Active MQ with")
	watchCommand.Flags().StringVarP(&password, 		"password", 	"p", 	os.Getenv("QNDMQ_PASSWORD"),	"The password to connect to Active MQ with")
	watchCommand.Flags().StringVarP(&queueName, 		"queue", 	"q", 	os.Getenv("QNDMQ_QUEUE"), 	"The name of the queue to post to")

	command.AddCommand(watchCommand)


	if err := command.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

