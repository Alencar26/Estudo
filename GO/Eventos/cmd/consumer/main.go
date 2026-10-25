package main

import (
	"fmt"

	"github.com/Alencar26/Estudo/tree/master/GO/Eventos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "minhafila")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) //informa que a msg foi lida e não precisa voltar para fila
	}
}
