package routes

import (
	"CRUD-Api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalidades).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadeById).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
