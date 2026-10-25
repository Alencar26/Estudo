package main

import (
	"log"
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
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
