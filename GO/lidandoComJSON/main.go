package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"` //tag
	Saldo  int `json:"s"` //tag
}

func main() {

	//ENCODER
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	//json sempre será retornado em bites (por isso a conversão para string)
	println(string(res))

	//retornando direto na saída padrão
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	//DECODER
	jsonPuro := []byte(`{"n": 2, "s": 200}`)

	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	validaErro(err)
	println(contaX.Saldo)
}

func validaErro(err error) {
	if err != nil {
		panic(err)
	}
}
