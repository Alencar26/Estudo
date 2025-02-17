package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	/*
		    Utilizando a solução do exercício anterior:

			1. Em package-level scope, atribua os seguintes valores às variáveis:
				1. para "x" atribua 42
				2. para "y" atribua "James Bond"
				3. para "z" atribua true
			2. Na função main:
				1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
				2. Demonstre a variável "s".
	*/

	//ANTES
	fmt.Println("TIPOS E VALORES ZEROS:")
	fmt.Printf("tipo: %T, valor: %v\n", x, x)
	fmt.Printf("tipo: %T, valor: %v\n", y, y)
	fmt.Printf("tipo: %T, valor: %v\n", z, z)

	//DEPOIS
	s := fmt.Sprint(x, y, z)
	fmt.Printf("tipo: %T, valor: %v\n", s, s)
}
