package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Alencar26/Estudo/GO/APIS/internal/dto"
	"github.com/Alencar26/Estudo/GO/APIS/internal/entity"
	"github.com/Alencar26/Estudo/GO/APIS/internal/infa/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary         Get a user JWT
// @Description     Get a user JWT
// @Tags            users
// @Accept          json
// @Produce         json
// @Param           request   body          dto.GetJWTInput true "user credentials"
// @Success         200       {object}      dto.GetJwtOutput
// @Failure         404       {object}      dto.Error
// @Failure         500       {object}      dto.Error
// @Router          /users/generate_token   [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpireIn := r.Context().Value("expireIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpireIn)).Unix(),
	})
	accessToken := dto.GetJwtOutput{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// create user godoc
// @summary       Create User
// @Description   Create User
// @Tags          users
// @Accept        json
// @Produce       json
// @Param         request   body      dto.CreateUserInput true "user request"
// @Success       201
// @Failure       500       {object}  dto.Error
// @Router        /users    [post]
func (h *UserHandler) CreateClient(w http.ResponseWriter, r *http.Request) {

	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Email, user.Name, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
