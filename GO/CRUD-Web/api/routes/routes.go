package routes

import (
	"CRUD-Web/internal/handlers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/new", handlers.New)
	http.HandleFunc("/insert", handlers.Insert)
	http.HandleFunc("/delete", handlers.Delete)
	http.HandleFunc("/edit", handlers.Edit)
	http.HandleFunc("/update", handlers.Update)
}
