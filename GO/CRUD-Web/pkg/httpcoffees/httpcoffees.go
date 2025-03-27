package httpcoffees

import (
	"io"
	"log"
	"net/http"
)

func GetCoffees() (body []byte, err error) {

	resp, err := http.Get("https://api.sampleapis.com/coffee/hot")
	if err != nil {
		log.Printf("Erro na requisição GET para https://api.sampleapis.com/coffee/hot: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler response body da requisição: %v\n", err)
		return nil, err
	}

	return body, nil
}
