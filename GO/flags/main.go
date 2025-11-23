package main

import (
	"flag"
	"fmt"
)

func main() {

	port := flag.Int("port", 8080, "Porta do servidor")
	debug := flag.Bool("debug", false, "habilita mode debug")
	name := flag.String("name", "André Alencar", "nome do usuário")
	//Flag de abreviações
	p := flag.Int("p", 8080, "porta (alias)")

	flag.Parse() // é necessário essa linha
	//Antes dessa chamada, as flags ainda não foram lidas da linha de comando.

	fmt.Println("Portal: ", *port)
	fmt.Println("Debug: ", *debug)
	fmt.Println("Name: ", *name)
	fmt.Println("p: ", p)

	/*
		EXEMPLO DE COMO RODAR

		go run main.go -> retorna os valores default
		go run main.go --port=3000 --debug=true --name=suco


	*/

	//Após o flag.Parse() tudo que sobrar será retornado pelo flag.Args()
	//É retornado uma lista com os valores passado "a mais"
	args := flag.Args()
	fmt.Println("Args: ", args)

	/*
		EXEMPLO

		go run main.go --port=3000 --debug=true --name=sudo esse_vai_sobrar esse_também
		go run main.go --port=9000 file1.txt file2.txt

	*/
}
