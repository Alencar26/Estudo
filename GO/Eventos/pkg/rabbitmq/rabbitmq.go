package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,         //nome da fila que quero consumir
		"go-consumer", //nome da aplicação que vai consumir
		false,         //autoAck - Quando a msg é lida o rabbitmq já dá baixa na msg e remove da fila
		false,         //fila exclusiva
		false,         //noLocal
		false,         //noWait
		nil,           //args amqp.Table
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, exchangeName string) error {
	err := ch.Publish(
		exchangeName, //Exchange (default do rabbitmq)
		"",           //Key
		false,        //Mandatório
		false,        //imediato
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
