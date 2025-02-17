package main

import "fmt"

// declaração variáveis global
var moto = "Honda CG"

func main() {

	fmt.Println("Olá mundo!")

	//declarçaão de variável dentro de uma bloco de código:
	carro := "Honda Fit"
	fmt.Println(carro)

	//usando variável global
	fmt.Println(moto)

	//Mais declarações

	//podes-se declarar duas variáveis ou mais assim:
	x, y := true, 34 //Não se pode declarar uma variável e não usar.
	fmt.Printf("%v, %v\n", x, y)

	//!ATENÇÃO
	// o simbolo (:=) serve apensas para declarar variáveis. (não para atribuir valores)
	//pode-se inferir o tipo da variável também.

	var z int = 33
	fmt.Println(z)
}
