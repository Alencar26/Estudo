package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
	http.HandleFunc("/", buscaCEP)
	http.ListenAndServe(":8080", nil)
}

func buscaCEP(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"mensagem": "Caminho URL não existe. Volte para raiz(/)."}`))
		return
	}

	cepParam := req.URL.Query().Get("cep")
	if cepParam == "" {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"mensagem": "Verifique se o CEP foi passado corretamente."}`))
		return
	}

	cep, err := getViaCEP(cepParam)
	if err != nil {
		log.Printf("Erro na função getViaCEP. Erro: %v", err)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"mensagem": "Ocorreu um erro inesperado.", "Error": "` + err.Error() + `"}`))
		return
	}

	cepJson, err := json.Marshal(cep)
	if err != nil {
		log.Printf("Erro parse Struct to Json. Erro: %v", err)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"mensagem": "Ocorreu um erro inesperado.", "Error": "` + err.Error() + `"}`))
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(cepJson))
}

func getViaCEP(cep string) (*CEP, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		log.Printf("Erro durante requisição para o ViaCEP. Erro: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao obter response da requisição. Erro: %v", err)
		return nil, err
	}

	var cepStruct CEP
	err = json.Unmarshal(body, &cepStruct)
	if err != nil {
		log.Printf("Erro ao fazer o parse da resposta: %v", err)
		return nil, err
	}

	return &cepStruct, nil
}
