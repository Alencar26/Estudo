package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Main é a Thread 1
func main() {

	//SEM GOROUTINES - Rodando de forma sequencial
	//task("A")
	//task("B")

	//COM GOROUTINES - Rodando de forma concorrente
	go task("A") //Thread 2
	go task("B") //Thread 3
	//Go Routine anonima
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "ANONIMA")
			time.Sleep(1 * time.Second)
		}

	}()

	time.Sleep(15 * time.Second) //se não tive o timer ou algo para segurar o processo principal o GO finaliza a execução.
}
