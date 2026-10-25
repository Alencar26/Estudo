package main

import "fmt"

// Validando valores zeros
var a int
var b float64
var c string
var d bool

func main() {

	//VALORES ZEROS
	// Valores padrão dos tipos de variáveis que ainda não tiveram uma inicialização/atribuição de valores.

	// - int: 0
	// - floats: 0.0
	// - booleans: false
	// - strings: ""
	// - pointers, functions, interfaces, slices, channels, maps: nil

	fmt.Println("\n[Valores Zeros] -----------------")

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Println("[Valores Zeros] ----------------- fim")

	//----------------------------------------
}
