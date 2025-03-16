package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	//USO DE TIMEOUT PARA REQUISIÇÕES EXTERNAS
	//É possível configurar nosso objeto Client antes da requisição.
	c := http.Client{Timeout: time.Second} //configurando tempo de timeout para requisição usando "c"
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() //fechando conexão no final de tudo

	body, err := io.ReadAll(resp.Body) //lendo o corpo da requisição
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
