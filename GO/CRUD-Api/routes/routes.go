package routes

import (
	"CRUD-Api/controllers"
	"CRUD-Api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalidades).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalidade).Methods(http.MethodPost)
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadeById).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidadeById).Methods(http.MethodDelete)
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonalidadeById).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
	//CORS permite que qualquer origem acesse a API
	// handlers.AllowedOrigins([]string{"*"}): Permite que qualquer origem acesse a API
	// handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}): Permite que qualquer método HTTP acesse a API
	// handlers.AllowedHeaders([]string{"Content-Type"}): Permite que qualquer cabeçalho HTTP acesse a API
	// handlers.AllowCredentials(): Permite que credenciais sejam enviadas na requisição
	// handlers.MaxAge(3600): Define o tempo máximo em segundos que o navegador deve armazenar as informações de CORS
	// handlers.ExposeHeaders([]string{"Content-Type"}): Permite que o cabeçalho Content-Type seja exposto na resposta
}
