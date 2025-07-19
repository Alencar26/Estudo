package main

import "fmt"

//ESSE CÓDIGO VAI GERAR UM DEADLOCK.
/*
No momento que escrever no channel a thread principal vai travar
esperando que outra thread leia o channel.

Como não há outra thread o Go retorna erro de deadlock.
*/
func main() {

	ch := make(chan string)

	ch <- "string aleatória"

	fmt.Println(<-ch)
}
