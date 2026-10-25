package main

func main() {

	//criando canal com buffer de 3 valores ao mesmo tempo no canal
	ch := make(chan string, 3)
	ch <- "valor 1"
	ch <- "valor 2"
	ch <- "valor 3"

	println(<-ch)
	println(<-ch)
	println(<-ch)
}
