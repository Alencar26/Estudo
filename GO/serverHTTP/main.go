package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("Roda não existe."))
		return
	}

	//Query Parameters. Ex: localhost:8080?nome=André
	nomeParam := req.URL.Query().Get("nome")
	if nomeParam == "" {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"mensagem":"Nome ausente.","status": "400 Bad Request"}`))
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"mensagem":"Olá ` + nomeParam + `!","status": "200 OK"}`))
}
