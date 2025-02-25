package main

import "fmt"

// usa-se ~ para que tipos custom possam ser compatíveis. (não obrigatório)
// se não usar, não é possível que tipos custom usem o generics
type Number interface {
	~int | ~float64
}

func soma[T Number](m map[string]T) T {
	var soma T
	for _, valor := range m {
		soma += valor
	}
	return soma
}

func main() {
	inteiro := map[string]int{"a": 12, "b": 55, "c": 30, "d": 6}
	flutuante := map[string]float64{"a": 12.7, "b": 5.89, "c": 30.3, "d": 6.5}

	fmt.Println(soma(inteiro))
	fmt.Println(soma(flutuante))
}
