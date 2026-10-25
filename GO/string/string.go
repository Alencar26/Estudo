package main

import (
	"fmt"
)

func main() {

	//String são imutáveis no GO
	s := "Hello playground"
	sb := []byte(s)
	//format string para utilizar no printf
	// %v, %T, %#U, %#x

	fmt.Printf("%v\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%#U\n", s)
	fmt.Printf("%#x\n", s)

	fmt.Printf("Slice de bytes da string, podemos ver o código UTF8 de cada caractere: %v\n", sb)

	fmt.Printf("%#x\n", sb[0])

	//range
	//pega o caractere em UTF8 (que pode ter mais de 1 byte)
	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	//for loop
	//pega byte a byte dos caracteres
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
