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

qndmq post Hello World!
```

## Credits

This project is made possible by the following awesome people:

+ [spf13](https://github.com/spf13) for [Cobra](https://github.com/spf13/cobra) released under the [Apache 2.0 license](https://github.com/spf13/cobra/blob/master/LICENSE.txt).
+ [vcabbage](https://github.com/vcabbage) for 1.0 compatible [amqp](https://github.com/vcabbage/amqp) released uder the [MIT license](https://github.com/vcabbage/amqp/blob/master/LICENSE)
+ [streadway](https://github.com/streadway) for 0.9.1 compatible [amqp](https://github.com/streadway/amqp) released under the [BSD license]()https://github.com/streadway/amqp/blob/master/LICENSE
