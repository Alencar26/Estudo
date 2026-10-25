package main

import "fmt"

func main() {
	hello := make(chan string)
	go recebe("Hello World", hello) //umas das funções precisa rodar em background em outra thread
	ler(hello)
}

// podemos deixar clano na assinatura da função que o canal só vai receber dados. Add <- na assinatura da função
func recebe(nome string, ch chan<- string) {
	ch <- nome
}

// podemos deixar claro na assinatura da função que o canal vai enviar os dados para alguém.
func ler(ch <-chan string) {
	fmt.Println(<-ch)
}
