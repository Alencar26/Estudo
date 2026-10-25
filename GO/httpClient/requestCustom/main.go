package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	client := http.Client{}

	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	//informando que eu aceito receber json (essa informação vai no header)
	req.Header.Add("Accept", "application/json")

	//agora meu client está usando a request custom que configuramos.
	resp, err := client.Do(req) //função para juntar nosso client com a request
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//jogando o retorno no meu terminal
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
