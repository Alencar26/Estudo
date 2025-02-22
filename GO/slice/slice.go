package main

import "fmt"

func main() {

	//slice trabalha com array por baixo do pano
	//utiliza-se de ponteiro para apontar para o array (possui tamnho e capacidade)

	//declarando um slice = igual array mas não fixa valor no []
	s := []int{2, 3, 1, 5, 6, 1, 5, 9}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//zerando o slice
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	//deixando só o quatro primeiros itens
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	//ignorando os dois primeiros e exiber o resto
	//a capacidade deminui pois estamos removendos os dois primeiros
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	//slice sem inicializar com valor
	d := []int{}
	d = append(d, 3)
	println(d[len(d)-1])

	//-------------------------

	carros := []string{}
	carros = append(carros, "gol")
	carros = append(carros, "honda")
	carros = append(carros, "palio")
	carros = append(carros, "ferrari")

	carros[2] = "FORD"

	for _, nome := range carros {
		println(nome)
	}
}
