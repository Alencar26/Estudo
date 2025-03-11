package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int) //criando canal de dados que será consimido pelo workers
	go worker(1, data)     //vai rodar em uma Thread independente
	go worker(2, data)     //vai rodar em uma Thread independente
	go worker(3, data)     //vai rodar em uma Thread independente

	//INICIANDO 50 WORKERS
	for w := range 50 {
		go worker(w, data)
	}

	//conforme o canal "data" vai recebendo informações, os workers vão executando de modo concorrente
	for i := range 1000 {
		data <- i
	}
}
