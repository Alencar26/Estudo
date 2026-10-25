package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}
 
  //buffer é um "espaço" cheio de informações
  jsonVar := bytes.NewBuffer([]byte(`{"name": "André"}`))

  resp, err := c.Post("http://google.com", "application/json", jsonVar)
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  io.CopyBuffer(os.Stdout, resp.Body, nil) //essa função pega o retorno de um lugar e joga na saída de outro lugar
  //ex: estou pegando o meu body e jogando no Stdout (terminal)
  }
