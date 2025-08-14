package main

import "github.com/Alencar26/Estudo/tree/master/GO/Eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
