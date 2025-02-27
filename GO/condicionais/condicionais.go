package main

import (
	"math/rand"
)

func main() {
	a := 1
	b := 2
	c := 3

	//no GO não existe elif, elseif...
	if a > b && c > a {
		println("END")
	} else if a == b {
		println(b)
	} else {
		println("d")
	}

	if b < a || c > b {
		println("OR")
	}

	x := rand.Intn(99)

	switch x % 5 {
	case 0:
		println("Multiplo de 5", x)
	case 1:
		println("Resta 1", x)
	case 2:
		println("Resta 2", x)
	case 3:
		println("Resta 3", x)
	case 4:
		println("Resta 4", x)
	default:
		println("Valor default")
	}
}
