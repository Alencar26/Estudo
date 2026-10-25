package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemmento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, ceps := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + ceps + "/json/")
		if err != nil {
			//Fprintf = vai jogar o resultado para os.Stderr
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			//log sempre manda a mensagem para o os.Stderr
			log.Printf("Erro ao ler response da requisição: %v\n", err)
		}

		var cep CEP
		err = json.Unmarshal(res, &cep)
		if err != nil {
			log.Printf("Erro ao fazer o parse da resposta: %v", err)
		}

		fmt.Println(cep)
		file, err := os.Create("DadosCEP.txt")
		if err != nil {
			log.Printf("Erro ao criar arquivo: %v", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", cep.Cep, cep.Localidade, cep.UF))
		if err != nil {
			log.Printf("Erro ao gravar arquivo: %v", err)
		}
	}
}
