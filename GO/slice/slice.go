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

	//juntando dois slices

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1[:])
	slice2 := []int{6, 7, 8, 9, 10}
	fmt.Println(slice2[:])

	slice1 = append(slice1, slice2...)
	fmt.Println(slice1[:])

	//------------------------------------------------------

	//usando make para criar os slices.
	//vantagem é que podemos já definir um tamanha maior no array por baixo dos panos
	//melhora a performance do software
	//exemplo: make([]T, len, cap) - len = tamanho inicial, cap = capacidade total do array por baixo.
	// se o len superar o cap então do GO vai criar outro array com 2x cap.

	x := make([]int, 5, 10)
	x[0], x[1], x[2], x[3], x[4] = 1, 2, 3, 4, 5 //len inicial 5
	x = append(x, 6)                             //posso adicionar + 5 elementos que que o array que esta por baixo mude para outro.
	fmt.Println(x[:])

	//------------------------------------PEGADINHA DO GO
	fmt.Println("\n-------------------------- PEGADINHA DO SLICE")
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1[:], "Slice 1")

	s2 := append(s1[:2], s1[4:]...)
	fmt.Println(s2[:], "Slice 2")

	fmt.Println(s1[:], "Slice 1")

	fmt.Println("O Array subjacente da Slice1 teve seu valor alterado indevidamente\npois o Slice2 faz o uso do mesmo array.")

	//Evite usar vários slices efetuando recortes de um slice para usar em outro
	//O GO pode acabar usando o mesmo array e modificando valores que não deveria
	//Nesse nosso caso a variável s1 teve seu valor alterado indevidamente.
	//Ou utilizae o mesmo slice e vátrabalhando com ele ou faça um for loop para
	//efetuar passagem dos valores para outro slice.
}
