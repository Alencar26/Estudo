package main

import (
	"fmt"
	"reflect"
)

type Pessoa struct {
	Nome  string
	Idade int
}

/*
O pacote reflect em Go permite inspeção e manipulação de tipos e valores em tempo de execução.
Ele fornece funcionalidades para obter o tipo (TypeOf), valor (ValueOf), e criar novas instâncias (New).
Além disso, permite acessar e modificar campos de estruturas, verificar igualdade profunda (DeepEqual), e obter valores zero de tipos (Zero).
É útil para metaprogramação e manipulação dinâmica de dados.
*/

func main() {

	var a string = "Oi"
	var b int = 1
	var c float64 = 6.4
	var d float32 = 3.2
	var e bool = true

	fmt.Println("O tipo da sua variável é:", reflect.TypeOf(a))
	fmt.Println("O tipo da sua variável é:", reflect.TypeOf(b))
	fmt.Println("O tipo da sua variável é:", reflect.TypeOf(c))
	fmt.Println("O tipo da sua variável é:", reflect.TypeOf(d))
	fmt.Println("O tipo da sua variável é:", reflect.TypeOf(e))
	println("")
	//-------------------------

	fmt.Println("O valor da variável é:", reflect.ValueOf(a))
	fmt.Println("O valor da variável é:", reflect.ValueOf(b))
	fmt.Println("O valor da variável é:", reflect.ValueOf(c))
	fmt.Println("O valor da variável é:", reflect.ValueOf(d))
	fmt.Println("O valor da variável é:", reflect.ValueOf(e))
	println("")

	//-------------------------

	fmt.Println("O valor zero de ua variável String é:", reflect.Zero(reflect.TypeOf(a)))
	fmt.Println("O valor zero de ua variável Int é:", reflect.Zero(reflect.TypeOf(b)))
	fmt.Println("O valor zero de ua variável Float64 é:", reflect.Zero(reflect.TypeOf(c)))
	fmt.Println("O valor zero de ua variável Float32 é:", reflect.Zero(reflect.TypeOf(d)))
	fmt.Println("O valor zero de ua variável Bool é:", reflect.Zero(reflect.TypeOf(e)))
	println("")

	//------------------------- EXEMPLO relcect.New

	//obtem o tipo passado no parâmetro
	t := reflect.TypeOf(Pessoa{})

	//cria nova instância a partir do tipo passado
	novaInstancia := reflect.New(t)

	//obtem o valor da instância
	valores := novaInstancia.Elem()

	//obtem o campo Nome
	valores.FieldByName("Nome").SetString("João")
	valores.FieldByName("Idade").SetInt(20)

	pessoaFinal := valores.Interface().(Pessoa)
	fmt.Println(pessoaFinal)
	println("")

	//-------------------------

	novaPessoa := CriarInstancia(Pessoa{})
	println(novaPessoa)
	println("")

	//-------------------------

	tipoEhIgual := reflect.DeepEqual(reflect.TypeOf(a), reflect.TypeOf(b))
	fmt.Println("Comparando os tipos: ", tipoEhIgual)
	println("")

	//-------------------------

}

func CriarInstancia(modelo interface{}) interface{} {
	t := reflect.TypeOf(modelo)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return reflect.New(t).Interface()
}
