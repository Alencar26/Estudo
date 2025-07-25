package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg) //mesmo essa função rodando em background o programação não fecha por causa do WaitGroup

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Printf("Recebido: %d\n", v)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
