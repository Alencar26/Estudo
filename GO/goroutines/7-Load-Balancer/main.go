package main

import (
	"fmt"
	"time"
)

func worker(worckerId int, datas <-chan int) {
	for x := range datas {
		fmt.Printf("Worker %d recebeu %d.\n", worckerId, x)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	data := make(chan int)

	workes := 1000000
	//inicializa os workes
	for i := 0; i < workes; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000000; i++ {
		data <- i
	}
}
