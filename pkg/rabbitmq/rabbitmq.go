package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel(stringConnection string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(stringConnection)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Producer(ch *amqp.Channel, body string, exchengeName string) error {
	err := ch.Publish(
		exchengeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
