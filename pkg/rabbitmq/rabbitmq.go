package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel(stringConnection string) (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:7070/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}
