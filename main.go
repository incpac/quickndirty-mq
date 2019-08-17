package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/incpac/quiet"
	"github.com/incpac/quiet/config"
	"github.com/spf13/cobra"
)

var Version string

var connectionString string
var username string
var password string
var queueName string


func createConnection() quiet.Client {
	conf := config.ParseString(connectionString)

	if username != "" {
		conf.Username = username 
	}

	if password != "" {
		conf.Password = password
	}

	if queueName != "" {
		conf.Queue = queueName
	}

	c, err := quiet.NewClient(conf) 
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func post(m string) {
	c := createConnection()

	c.Post(m)
	c.Close()
}


func watch() {
	c := createConnection()
	
	c.Watch(func(s string) {
		log.Printf("Message received: %s", s)
	})

	log.Println("Watching queue...")

	// run forever
	for {}

	c.Close()
}

func get() {
	c := createConnection()

	msg, err := c.Get()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Retrieved message: %s", msg)
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
		Run:   func(cmd *cobra.Command, args []string) {
			post(strings.Join(args, " "))
		},
	}

	postCommand.Flags().StringVarP(&connectionString,	"connection",	"c",	os.Getenv("QNDMQ_CONNECTION"),	"The connection string for the MQ server")
	postCommand.Flags().StringVarP(&username,		"username",	"u",	os.Getenv("QNDMQ_USERNAME"),	"The username to connect to MQ with")
	postCommand.Flags().StringVarP(&password,		"password",	"p",	os.Getenv("QNDMQ_PASSWORD"),	"The password to connect to MQ with")
	postCommand.Flags().StringVarP(&queueName,		"queue",	"q",	os.Getenv("QNDMQ_QUEUE"),	"The name of the queue to post to")

	command.AddCommand(postCommand)


	watchCommand := &cobra.Command{
		Use:   "watch",
		Short: "Watch the queue for new messages",
		Long:  "Watch the queue for new messages",
		Run:   func(cmd *cobra.Command, args []string) {
			watch()
		},
	}

	watchCommand.Flags().StringVarP(&connectionString,	"connection",	"c",	os.Getenv("QNDMQ_CONNECTION"),	"The connection string for the MQ server")
	watchCommand.Flags().StringVarP(&username,		"username",	"u",	os.Getenv("QNDMQ_USERNAME"),	"The username to connect to MQ with")
	watchCommand.Flags().StringVarP(&password,		"password",	"p",	os.Getenv("QNDMQ_PASSWORD"),	"The password to connect to MQ with")
	watchCommand.Flags().StringVarP(&queueName,		"queue",	"q",	os.Getenv("QNDMQ_QUEUE"),	"The name of the queue to post to")

	command.AddCommand(watchCommand)


	getCommand := &cobra.Command{
		Use:	"get",
		Short:	"Get the next message in the queue",
		Long:	"Get the next message in the queue",
		Run:	func(cmd *cobra.Command, args []string) {
			get()
		},
	}

	getCommand.Flags().StringVarP(&connectionString,	"connection",	"c",	os.Getenv("QNDMQ_CONNECTION"),	"The connection string for the MQ server")
	getCommand.Flags().StringVarP(&username,		"username",	"u",	os.Getenv("QNDMQ_USERNAME"),	"The username to connect to MQ with")
	getCommand.Flags().StringVarP(&password,		"password",	"p",	os.Getenv("QNDMQ_PASSWORD"),	"The password to connect to MQ with")
	getCommand.Flags().StringVarP(&queueName,		"queue",	"q",	os.Getenv("QNDMQ_QUEUE"),	"The name of the queue to post to")

	command.AddCommand(getCommand)


	if err := command.Execute(); err != nil {
		log.Fatal("Failed to start:", err)
		os.Exit(-1)
	}
}

