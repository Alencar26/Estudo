package main

import "fmt"

func main() {

	x := 1 //bit 001
	//deslocando bit para esquerda
	y := x << 2 //bit 100
	z := y >> 1 //bit 010

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
	fmt.Println("----------------------")

	//exemplo de uso
	const (
		_  = iota             //0
		KB = 1 << (iota * 10) // 1 << (1 * 10)
		MB                    // 1 << (2 * 10)
		GB                    // 1 << (3 * 10)
		TB                    // 1 << (4 * 10)
	)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
