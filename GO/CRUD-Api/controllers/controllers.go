package controllers

import (
	"CRUD-Api/database"
	"CRUD-Api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page.")
}

func GetAllPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func GetPersonalidadeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CreatePersonalidade(w http.ResponseWriter, r *http.Request) {
	var newPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&newPersonalidade)
	database.DB.Create(&newPersonalidade)
	json.NewEncoder(w).Encode(newPersonalidade)
}

func DeletePersonalidadeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode("Personalidade deletada com sucesso.")
}

// func EditPersonalidadeById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	var personalidade models.Personalidade
// 	database.DB.First(&personalidade, id)

// 	var updatePersonalidade models.Personalidade
// 	json.NewDecoder(r.Body).Decode(&updatePersonalidade)

// 	database.DB.Model(&personalidade).Updates(updatePersonalidade)
// 	json.NewEncoder(w).Encode("Personalidade atualizada com sucesso.")
// }

func EditPersonalidadeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode("Personalidade atualizada com sucesso.")
}
