package main

import "fmt"

type pokemon int

var pikachu pokemon
var charizard pokemon

func main() {
	/*
			Utilizando a solução do exercício anterior:

		    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
		    2. Na função main:
		        1. Isto já deve estar feito:
		            1. Demonstre o valor da variável "x"
		            2. Demonstre o tipo da variável "x"
		            3. Atribua 42 à variável "x" utilizando o operador "="
		            4. Demonstre o valor da variável "x"
		        2. Agora faça tambem:
		            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
		            2. Demonstre o valor de "y"
		            3. Demonstre o tipo de "y"
	*/

	//ANTES
	pikachu = 101
	fmt.Printf("tipo: %T, valor: %v\n", pikachu, pikachu)

	//DEPOIS
	fmt.Println("Valor:", charizard)
	fmt.Printf("Tipo: %T\n", charizard)
	charizard = 42
	fmt.Println("Valor:", charizard)

	x := 5 //int
	fmt.Println("Valor:", x)
	fmt.Printf("Tipo: %T\n", x)

	charizard = pokemon(x)

	fmt.Printf("%v", charizard)
}
