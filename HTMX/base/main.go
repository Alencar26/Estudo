package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	t := template.Must(template.New("index.html").ParseFiles("public/index.html"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	})

	r.Get("/carrega-conteudo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h3>Carregou Conteúdo!</h3>"))
	})

	r.Post("/enviar-formulario", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			panic(err)
		}
		name := r.FormValue("name")
		fmt.Fprintf(w, "<h3>Formulário enviado. Nome %s</h3>", name)
	})

	r.Get("/trocar-conteudo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h3>Troquei o Conteúdo!</h3>"))
	})

	r.Post("/valores-adicionais", func(w http.ResponseWriter, r *http.Request) {
		//ParseForm para popular o r.Form com os valores do formulário
		//Ele monta o map[string][]string com os valores enviados
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		valorExtra := r.FormValue("valorExtra")
		user := r.FormValue("user")

		fmt.Fprintf(w, "<h3>Troquei o Conteúdo!</h3><br>Valor Extra: %s<br>User:%s", valorExtra, user)
	})

	r.Get("/tempo-servidor", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format(time.RFC850)
		fmt.Fprintf(w, "<h3>Hora atual do servidor: %s</h3>", now)
	})

	r.Post("/enviar-formulario-validacao", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			panic(err)
		}
		username := r.FormValue("username")
		if username == "" || len(username) < 6 {
			fmt.Fprint(w, "<p style='color: red;'>Error: O nome de usuário deve ter pelo menos 6 caracteres.</p>")
		} else {
			fmt.Fprintf(w, "<p>Formulário enviado com sucesso. Nome de usuário: %s</p>", username)
		}
	})

	http.ListenAndServe(":3000", r)
}
