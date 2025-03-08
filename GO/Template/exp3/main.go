package main

import (
	"log"
	"os"
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

	//poderia passar vários arquivos: template.New("template.html").ParseFiles([]string["template.html..."])
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"GO", 40},
		{"Java", 70},
		{"Python", 50},
		{"PHP", 80},
	})
	check(err)
}
