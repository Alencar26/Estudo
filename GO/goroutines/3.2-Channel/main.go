package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- "String enviada para o canal"
	fmt.Println("FIM")
}
