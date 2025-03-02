package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Olá Mundo!"))
}
