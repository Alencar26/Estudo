package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx := context.Background() //contexto vazio

	//qlqr requisição que estiver usando o contexto abaixo terá 1 segundo de timeout
	ctx, cancel := context.WithTimeout(ctx, time.Second) //adicionando um contexto de timeout de 1 segundo
	defer cancel()                                       //cancelando o contexto - para a execução do contexto da request

	//criando uma request com contexto
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	//fazendo a requisição, juntnado com client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
