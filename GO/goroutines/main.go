package main

import (
	"fmt"
	"time"
)

func contador(n int) {
	for i := range n {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	contador(10) //assim o código vai ser ligo de sima para baixo sequencialmente
	contador(5)  //assim o código vai ser ligo de sima para baixo sequencialmente

	//Agora eu consigo paralelizar a execução
	go contador(5) //vai executar em paralelo com a função abaixo
	contador(10)
}
