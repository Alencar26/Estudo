package main

import "fmt"

func main() {

	//Decimal		%d
	//Hexadecimal	%#x
	//Unicode		%#U
	//Tab			\t
	for i := 33; i <= 122; i++ {
		fmt.Printf("O número %d\t significa = %s\n", i, string(i))
	}
}
