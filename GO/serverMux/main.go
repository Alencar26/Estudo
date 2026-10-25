package main

import "net/http"

func main() {
	//multplex GO
	mux := http.NewServeMux() //mux personalizado
	mux.Handle("/blog", blog{title: "Meu blog"})
	http.ListenAndServe(":8080", mux) //server com mux custom

	//é possível ter mais de um server rodando no mesmo binário GO usando MUX
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/blog", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Esse é um outro server rodando. Não tem problema ter o mesmo endpoint do outro server."))
	})
	http.ListenAndServe(":8081", mux2) //usando o mux 2 (outro server)
}

//Essa forma dá muita flexibilidade para customizar o meu handler.

type blog struct {
	title string
}

func (b blog) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(b.title))
}
