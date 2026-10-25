package main

import "time"

//SOME OS PRIMEIROS NÚMEROS NATURAIS ATÉ N

// menos performático
func somarAteN(n int) (soma int) {
	for n > 0 {
		soma += n
		n--
	}
	return
}

// MAIS PERFORMÁTICO
func somarAteN2(n int) int {
	return n * (n + 1) / 2
}

func main() {

	timestart := time.Now()
	println(somarAteN(999999999))
	timeend := time.Now()

	println("O tempo da 1º função foi de:", timeend.Sub(timestart).Milliseconds(), "Milissegundos.")

	timestart = time.Now()
	println(somarAteN2(999999999))
	timeend = time.Now()
	println("O tempo da 2º função foi de:", timeend.Sub(timestart).Milliseconds(), "Milissegundos.")
}
