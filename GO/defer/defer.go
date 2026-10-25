package main

import (
	"io"
	"net/http"
)

func main() {

	//EXEMPLO FUNCIONAMENTO DEFER
	funcionamentoDefer()

	//requestHTTP()
}

func funcionamentoDefer() {
	println("Primeira linha")
	defer println("Segunda linha") //vai ser executado no final (segura execução)
	println("Terceira linha")
	println("Quarta linha")
	println("Quinta linha")
}

func requestHTTP() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	//quando começa com defer, significa que só será executado quando todas as linhas do arquivo
	//já tiverem sido executadas; (É a última coisa que será executada - há um atraso na execução)
	defer request.Body.Close()

	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
