package main

import "fmt"

func main() {

	var i uint16
	i = 65535
	fmt.Println(i)

	//Aqui vai estourar erro pois o máximo é 65535
	//i = 65536
	//fmt.Println(i)

	//Agora aqui, ele vai funcionar igual um hodômetro de carro, uando chega
	//no último valor ele zera e começa a contagem.
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
}
