package main

import "fmt"

//como duas Goroutines (threads) conversam e compartilham informações?
//R: usando canais de comunicação.

//O GO contorna o problema de condição de corrida usando canais de comunicação
//Assim evita que uma mesma variável seja manipulada ao mesmo tempo por duas Threads diferentes.

// Thread 1
func main() {

	channel := make(chan string) //criando um canal de string

	//Thread 2
	go func() {
		channel <- "Hello World." //enviando mensagem para o canal
	}()

	msg := <-channel //recebe mensagem do canal
	fmt.Println(msg)
}
