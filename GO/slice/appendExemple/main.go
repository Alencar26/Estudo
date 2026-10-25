package main

import "fmt"

func main() {

	evento := []string{"teste1", "teste2", "teste3", "teste4"}
	evento1 := append(evento[:0], evento[1:]...)
	fmt.Println(evento1)
	fmt.Println(evento[:0])
}
