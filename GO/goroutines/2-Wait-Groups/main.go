package main

import (
	"fmt"
	"sync"
	"time"
)

// passa um ponteiro, pois eu quero pegar um waitGroup específio. Senão passar ponteiro o Go vai passar uma cópia.
func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Main é a Thread 1
func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(27) //vai aguardar 27 operações
	//Thread 2
	go task("A", &waitGroup)
	//Thread 3
	go task("B", &waitGroup)
	//Thread 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "ANONIMA")
			time.Sleep(1 * time.Second)
			waitGroup.Done() //não precisa passar por parâmentro pois está dentro do contexto da variável waitGroup
		}
	}()
	waitGroup.Wait()
}
