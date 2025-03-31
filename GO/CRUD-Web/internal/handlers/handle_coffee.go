package handlers

import (
	"CRUD-Web/db"
	"CRUD-Web/internal/entities"
	"CRUD-Web/internal/repositories"
	"CRUD-Web/internal/services"
	"log"
	"net/http"
	"strconv"
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
		log.Printf("Handler Index: Falha ao chamar GetAllCoffees(): %V\n", err)
		panic(err.Error())
	}

	tmpl.ExecuteTemplate(w, "Index", coffees)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		// quantidade := r.FormValue("quantidade")

		precoToFloat64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Printf("Falha na conversão do preco (string) para float64: %V\n", err)
			panic(err.Error())
		}

		db := db.ConnectDB()
		defer db.Close()

		repo := repositories.NewCoffeeRepository(db)
		service := services.NewCoffeeService(repo)

		coffee := entities.Coffee{
			Title:       nome,
			Description: descricao,
			Price:       precoToFloat64,
			Ingredients: entities.Ingredients{
				"Água",
				"Café",
			},
			Image: "NULL",
		}
		service.CreateCoffee(&coffee)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Erro ao converter id (string) para inteiro: %V\n", err)
	}

	db := db.ConnectDB()
	defer db.Close()

	repo := repositories.NewCoffeeRepository(db)
	service := services.NewCoffeeService(repo)

	service.DeleteCoffee(idToInt)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Erro ao convertero ID string para int: %V", err)
		panic(err.Error())
	}

	var coffee *entities.Coffee

	db := db.ConnectDB()
	repo := repositories.NewCoffeeRepository(db)
	service := services.NewCoffeeService(repo)

	coffee, err = service.GetCoffeeById(idToInt)
	if err != nil {
		log.Printf("Erro ao obter o café pelo ID em GetCoffeeById(): %V", err)
		panic(err.Error())
	}

	tmpl.ExecuteTemplate(w, "Edit", coffee)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		// quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Falha na conversão do ID (string) para int: %V\n", err)
			panic(err.Error())
		}

		precoToFloat64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Printf("Falha na conversão do preco (string) para float64: %V\n", err)
			panic(err.Error())
		}

		db := db.ConnectDB()
		defer db.Close()

		repo := repositories.NewCoffeeRepository(db)
		service := services.NewCoffeeService(repo)

		coffee := entities.Coffee{
			Id:          idInt,
			Title:       nome,
			Description: descricao,
			Price:       precoToFloat64,
		}
		service.UpdateCoffee(&coffee)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
