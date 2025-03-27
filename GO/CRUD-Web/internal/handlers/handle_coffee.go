package handlers

import (
	"CRUD-Web/db"
	"CRUD-Web/internal/repositories"
	"CRUD-Web/internal/services"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	defer db.Close()

	repo := repositories.NewCoffeeRepository(db)
	service := services.NewCoffeeService(repo)

	coffees, err := service.GetAllCoffees()
	if err != nil {
		log.Printf("Handler Index: Falha ao chamar GetAllCoffees(): %v\n", err)
		panic(err.Error())
	}

	tmpl.ExecuteTemplate(w, "Index", coffees)
}
