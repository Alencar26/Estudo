package main

func main() {

	//só existe FOR no GO

	//for padrão
	for i := 0; i <= 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três", "quatro"}

	//for com range
	for k, v := range numeros {
		println(k, v)
	}

	//variação do for com range
	for _, v := range numeros {
		println("Sem mostrar chave:", v)
	}

	//for no estilo while (enquanto isso acontecer faça isso...)
	j := 0
	for j <= 10 {
		println(j)
		j++
	}

	controle := 0
	//loop infinito
	for {
		if controle == 10 {
			break
		}
		println("Infinito...")
		controle++
	}
}
