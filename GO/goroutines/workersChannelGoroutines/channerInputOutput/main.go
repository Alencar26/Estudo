package main

import (
	"fmt"
	"time"
)

func duplicateWorker(workerID int, input chan int, out chan<- int) {
	for x := range input {
		aux := 0
		fmt.Printf("Worker %d received %d\n", workerID, x)
		aux = x * 2
		out <- aux
		fmt.Printf("Worker %d output is %d\n", workerID, aux)
		time.Sleep(time.Second)
	}
}

func main() {

	//criando canais de entrada e saída
	chInput := make(chan int)
	chOutput := make(chan int)

	//criando workers para dobrar os valores e transferir de uma panal para outro
	for w := range 100 {
		go duplicateWorker(w, chInput, chOutput)
	}

	go func() {
		//populando canal de entrada
		for x := range 1000 {
			chInput <- x
		}
	}()

	i := 0
	for i <= 1000 {
		out := <-chOutput
		fmt.Printf("Result: %d\n", out)
		i++
	}
}
