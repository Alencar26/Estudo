package main

import "fmt"

func main() {

	//Posso usar qualquer tipo com interface vazia
	var a interface{} = "String"
	var b interface{} = 10
	var c interface{} = true
	println(a, b, c)

	//posso pegar uma variável de interface vazia e converter para o tipo que eu quero
	var x interface{} = "Andre"
	println(x.(string)) //convertendo para string

	result, ok := x.(int) //convertento para inteito
	fmt.Printf("O valor é %v e o resultado de ok é %v", result, ok)
	//O valor é 0 e o resultado de ok é false
}
