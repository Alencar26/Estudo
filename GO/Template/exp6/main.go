package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func TudoMaiusculo(s string) string {
	return strings.ToUpper(s)
}

func main() {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := template.New("content.html")
		t.Funcs(template.FuncMap{"TudoMaiusculo": TudoMaiusculo}) //mapeio minha funcção custom em um map
		t = template.Must(t.ParseFiles(templates...))

		err := t.Execute(w, Cursos{
			{"GO", 40},
			{"Java", 70},
			{"Python", 50},
			{"PHP", 80},
		})
		check(err)
	})
	http.ListenAndServe(":8080", nil)
}
