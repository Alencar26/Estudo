package routes

import (
	"CRUD-Web/internal/handlers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", handlers.Index)
}
