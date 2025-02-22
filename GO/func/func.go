package main

import "errors"

func main() {

	x := soma(2, 3)
	println(x)

	println(doisRetornos("Oi", "OOi"))
	println(doisRetornos("Oi", "Oi"))

	println(somaAte50(12, 12))
	println(somaAte50(25, 27))

	valor, err := somaAte50(25, 42)
	if err != nil {
		println(err.Error())
	}
	println(valor)
}

func soma(a int, b int) int {
	return a + b
}

func void(a, b int) {
	println("Não tem retorno", a, b)
}

func doisRetornos(a, b string) (int, bool) {
	if a != b {
		return 0, false
	} else {
		return 1, true
	}
}

func somaAte50(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma pasosu de 50")
	}
	return a + b, nil
}
