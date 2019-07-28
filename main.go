package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"pack.ag/amqp"
)

var Version string

var connectionString string
var username string
var password string
var queueName string


func post(cmd *cobra.Command, args []string) {

	_, err := amqp.Dial(connectionString, amqp.ConnSASLPlain(username, password))
	if err != nil {
		if err.Error() == "unexpected protocol version 0.9.1" {
			ohpointninePost(cmd, args)
		} else {
			log.Fatal(err)
		}
	} else {
		onepointohPost(cmd, args)
	}
}


func watch(cmd *cobra.Command, args []string) {

	_, err := amqp.Dial(connectionString, amqp.ConnSASLPlain(username, password))
	if err != nil {
		if err.Error() == "unexpected protocol version 0.9.1" {
			ohpointnineWatch(cmd, args)
		} else {
			log.Fatal("Failed to determine protocol:", err)
		}
	} else {
		onepointohWatch(cmd, args)
	}
}


func main() {

	command := &cobra.Command{
		Use:   "qndmq",
		Short: "A quick 'n' dirty Apache Active MQ client",
		Long:  "A simple Apache Active MQ client useful for testing configurations and running servers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", Version)
		},
	}

	
	postCommand := &cobra.Command{
		Use:   "post [flags] [message]",
		Short: "Post a message to the queue",
		Long:  "Post a message to the queue",
		Args:  cobra.MinimumNArgs(1),
		Run:   post,
	}

	postCommand.Flags().StringVarP(&connectionString,	"connection",	"c",	os.Getenv("QNDMQ_CONNECTION"),	"The connection string for the Active MQ server")
	postCommand.Flags().StringVarP(&username,		"username",	"u",	os.Getenv("QNDMQ_USERNAME"),	"The username to connect to Active MQ with")
	postCommand.Flags().StringVarP(&password,		"password",	"p",	os.Getenv("QNDMQ_PASSWORD"),	"The password to connect to Active MQ with")
	postCommand.Flags().StringVarP(&queueName,		"queue",	"q",	os.Getenv("QNDMQ_QUEUE"),	"The name of the queue to post to")

	command.AddCommand(postCommand)


	watchCommand := &cobra.Command{
		Use:   "watch",
		Short: "Watch the queue for new messages",
		Long:  "Watch the queue for new messages",
		Run:   watch,
	}

	watchCommand.Flags().StringVarP(&connectionString,	"connection",	"c",	os.Getenv("QNDMQ_CONNECTION"),	"The connection string for the Active MQ server")
	watchCommand.Flags().StringVarP(&username,		"username",	"u",	os.Getenv("QNDMQ_USERNAME"),	"The username to connect to Active MQ with")
	watchCommand.Flags().StringVarP(&password,		"password",	"p",	os.Getenv("QNDMQ_PASSWORD"),	"The password to connect to Active MQ with")
	watchCommand.Flags().StringVarP(&queueName,		"queue",	"q",	os.Getenv("QNDMQ_QUEUE"),	"The name of the queue to post to")

	command.AddCommand(watchCommand)


	if err := command.Execute(); err != nil {
		log.Fatal("Failed to start:", err)
		os.Exit(-1)
	}
}

