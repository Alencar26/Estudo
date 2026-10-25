package main

// significa que os valores de parâmetros passados são cópias dos endereços originais
func somaSemPonteiro(a, b int) int {
	a = 80
	return a + b
}

// significa que os valores de parâmetros passados são os originais na memória
func somaComPonteiro(a, b *int) int {
	*a = 80
	return *a + *b
}

func main() {

	//Memória -> Endereço -> Valor
	a := 10

	println(a)  //valor
	println(&a) //endereço na memória

	//quando colocar * significa que estou eferenciando o endereço
	var ponteiro *int = &a
	println(ponteiro)

	*ponteiro = 20
	println(a) //agora "a" vale 20 (mesmo endereço na memória)

	b := &a
	println(b)  //retorna o endereço
	println(*b) //retorna o valor desse endereço

	z := 20
	y := 30

	r1 := somaSemPonteiro(z, y)
	println(r1, "Sem Ponteiros -----------------------")
	println(z) //não teve alteração para 80
	println(y)

	r2 := somaComPonteiro(&z, &y)
	println(r2, "Com Ponteiros -----------------------")
	println(z) //sofreu alteração para 80
	println(y)

}
