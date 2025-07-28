package main

import (
	"fmt"
	"time"
)

type Message struct {
	id  int
	msg string
}

func main() {

	c1 := make(chan Message)
	c2 := make(chan Message)

	//RabbitMQ
	go func() {
		i := 0
		for {
			i++
			time.Sleep(time.Second * 2)
			c1 <- Message{i, "Message RabbitMQ"}
		}
	}()

	//Kafka
	go func() {
		i := 212
		for {
			i++
			time.Sleep(time.Second * 1)
			c2 <- Message{i, "Message Kafka"}
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout")
			//default:
			//	fmt.Println("Default")
		}
	}
}
