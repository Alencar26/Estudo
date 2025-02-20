package main

import "fmt"

func main() {

	var meuArray [3]int
	meuArray[0] = 43
	meuArray[1] = 22
	meuArray[2] = 66
	fmt.Printf("%v\n", meuArray[1])
	fmt.Printf("%v\n", meuArray)

	fmt.Println("tamanho do meu array:", len(meuArray))
	fmt.Println("Ultimo elemento:", meuArray[len(meuArray)-1])

	println("Percorrendo meu Array:")

	for i, v := range meuArray {
		println("Índice", i, "Valor:", v)
	}

	println("------------")

	for _, v := range meuArray {
		println("Só valor valor:", v)
	}

	println("------------")

	for i, _ := range meuArray {
		println("Só valor índice:", i)
	}

	println("------------")

	for i := range meuArray {
		println("Só valor índice:", i)
	}

}
