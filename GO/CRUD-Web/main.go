package main

import (
	"CRUD-Web/api/routes"
	"log"
	"net/http"
)

func main() {
	routes.LoadRoutes()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Falha no start do servidor: %v\n", err)
		panic(err)
	}
}
