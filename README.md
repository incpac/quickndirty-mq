# QuickNDirty-MQ

QuickNDirty-MQ is a stupidly simple [Apache ActiveMQ](https://activemq.apache.org/) client designed for testing connections to servers.  

This should work on all servers implementing the AMQP protocol, such as [Amazon MQ](https://aws.amazon.com/amazon-mq/).

## Compile

Simply download, grab dependencies, and compile.

```bash
go get github.com/incpac/quickndirty-mq
dep ensure
go build -ldflags "-X main.Version=0.1" -o qndmq
```

## Running

For a full list of commands and flags see `qndmq help`

### Watching the queue

```bash
qndmq watch -connection amqp://myamqpserver.example.com -u username -p password -q queue
```

### Pushing a message to the queue

```bash
qndmq post -connection amqp://myamqpserver.example.com -u username -p password -q queue Hello World!
```

### Environment variables

An alternative to passing through configuration via flags is to preset it as environment variables.

```bash
export QNDMQ_CONNECTION=amqp://myamqpserver.example.com
export QNDMQ_USERNAME=username
export QNDMQ_PASSWORD=password
export QNDMQ_QUEUE=queue

/qndmq post Hello World!
```
