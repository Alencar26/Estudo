package main

import "fmt"

func main() {

	//IOTA = adiciona valiores inteiros sequenciais não tipados como valor da constante
	// É possível ignoras os valores usando _
	const (
		x = iota
		_ = iota
		z = iota
		_ = iota
		b = iota
	)

	fmt.Println(x, z, b)

	//É possível adicioanr valor só na primeira const. O resto vai por tabela
	const (
		c = iota
		d
		e
		f
		g
	)

	fmt.Println(c, d, e, f, g)

	// É possível adicionar formula e todos refletiram a formula
	const (
		h = iota * 10
		i
		j
		k
		l
	)

	fmt.Println(h, i, j, k, l)
}
