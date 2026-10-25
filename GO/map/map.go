package main

import "fmt"

func main() {

	salario := map[string]int{"Andre": 1000, "João": 2000, "Pedro": 1500}

	fmt.Println(salario["Pedro"])

	frutas := map[string]string{
		"Vermelho": "Maça",
		"Amarelo":  "Banana",
		"Verde":    "Melancia",
		"Roxo":     "Uva",
	}

	//deletar valor
	delete(frutas, "Amarelo")

	//Adicionar valor
	frutas["Preta"] = "Jabuticaba"

	//Imprimir valor
	fmt.Println(frutas["Preta"])

	//Formas de criar map vazio
	carros := make(map[string]string)
	motots := map[string]string{}

	fmt.Println(carros)
	fmt.Println(motots)

	//percorrendo map
	for cor, fruta := range frutas {
		fmt.Println(cor+":", fruta)
	}

	//-------------------------

	//validando se valores existem no map
	fruta, ok := frutas["Azul"]
	if !ok {
		println("Fruta não existe.")
	} else {
		println(fruta)
	}

	//forma alternativa
	if fruta, ok = frutas["Laranja"]; !ok {
		println("Fruta não existe.")
	} else {
		println(fruta)
	}
}
