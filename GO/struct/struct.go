package main

import "fmt"

type Caracteristica struct {
	Cor    string
	Altura float64
	Peso   float64
	Genero string
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	UF         string
}

type Cliente struct {
	Nome      string
	Idade     int
	Ativo     bool
	Aparencia Caracteristica
	Endereco
}

// FUNÇÃO DA STRUCT "Método de Classe"
// * indica que estou usando um ponteiro para instância atual
func (c *Cliente) Desativar() {
	c.Ativo = false
	println(c.Ativo)
}

func main() {
	cliente1 := Cliente{
		Nome:  "André",
		Idade: 28,
		Ativo: true,
	}

	cliente1.Logradouro = "Rua dos bobos"
	cliente1.Numero = 0
	cliente1.Cidade = "Bobolândia"
	cliente1.UF = "Birutas"

	cliente1.Aparencia.Altura = 1.75
	cliente1.Aparencia.Peso = 80.0
	cliente1.Aparencia.Cor = "Verde"
	cliente1.Aparencia.Genero = "M"

	//chamando a função da struct "Método de classe"
	cliente1.Desativar()
	println(cliente1.Ativo)

	fmt.Printf("Nome: %s, Idade %d, Ativo: %t\n", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
	fmt.Printf(
		"Nome: %s, Idade %d, Ativo: %t, Rua: %s nº%d, Cidade: %s, Estado: %s\n",
		cliente1.Nome,
		cliente1.Idade,
		cliente1.Ativo,
		cliente1.Logradouro,
		cliente1.Numero,
		cliente1.Cidade,
		cliente1.UF,
	)
}
