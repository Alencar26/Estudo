package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	b = 30
	fmt.Printf("%T\n", b)
	//RETORNO: main.hotdog

	x := 10
	fmt.Printf("%T\n", x)
	//RETORNO: int

	//b = x
	//ISSO VAI DAR ERRO - NÃO SÃO DO MESMO TIPO

	b = hotdog(x)
	//AGORA SIM - fizemos a conversão para o tipo certo.

	fmt.Println("b:", b)
}
