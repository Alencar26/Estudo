package main

import "fmt"

type pokemon int

var pikachu pokemon

func main() {
	/*
			Crie um tipo. O tipo subjacente deve ser int.
		    Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.

			Na função main:
		    1. Demonstre o valor da variável "x"
		    2. Demonstre o tipo da variável "x"
		    3. Atribua 42 à variável "x" utilizando o operador "="
		    4. Demonstre o valor da variável "x"
	*/

	pikachu = 42
	fmt.Printf("tipo: %T, valor: %v\n", pikachu, pikachu)
}
