package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Alencar26/Estudo/GO/APIS/internal/dto"
	"github.com/Alencar26/Estudo/GO/APIS/internal/entity"
	"github.com/Alencar26/Estudo/GO/APIS/internal/infa/database"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary       Create product
// @Description   Criação de regstro de produtos na base de dados.
// @Tags          products
// @Accept        json
// @Produce       json
// @Param         request           body       dto.CreateProductInput true "product request"
// @Success       201
// @Failure       501               {object}   dto.Error
// @Router        /products         [post]
// @Security      ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary       Get product
// @Description   Listagem de todos os registros de produtos da base de dados.
// @Tags          products
// @Produce       json
// @Param         page              query      string                     false   "page number"
// @Param         limit             query      string                     false   "limit"
// @Param         sort              query      string                     false   "sort"
// @Success       200               {array}   dto.GetProductOutput
// @Failure       404               {object}   dto.Error
// @Failure       500               {object}   dto.Error
// @Router        /products         [get]
// @Security      ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Get Product by ID godoc
// @Summary       Get product
// @Description   Retorna um registro de produtos da base de dados utilizando-se do ID para buscar.
// @Tags          products
// @Produce       json
// @Param         id                path        string                     true   "id"  Format(uuid)
// @Success       200               {object}    dto.GetProductOutput
// @Failure       400               {object}    dto.Error
// @Failure       404               {object}    dto.Error
// @Failure       500               {object}    dto.Error
// @Router        /products/{id}         [get]
// @Security      ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := errors.New("ID inválido! Deve ser diferente de vazio.")
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary       Update product
// @Description   Atualização do restro de produto na base de dados utilizandos-se do ID para identificar o registro.
// @Tags          products
// @Accept        json
// @Produce       json
// @Param         id                path        string                      true   "id" Format(uuid)
// @Param         request           body        dto.CreateProductInput      true   "request update product"
// @Success       200
// @Failure       400               {object}    dto.Error
// @Failure       404               {object}    dto.Error
// @Failure       500               {object}    dto.Error
// @Router        /products/{id}    [put]
// @Security      ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := errors.New("ID inválido! Deve ser diferente de vazio.")
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	var productUpdate dto.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&productUpdate); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	product.Name = productUpdate.Name
	product.Price = productUpdate.Price

	err = h.ProductDB.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary       Delete product
// @Description   Deleção de um restro de produto na base de dados utilizando-se do ID para identificar o registro.
// @Tags          products
// @Produce       json
// @Param         id                path        string                      true   "id" Format(uuid)
// @Success       200
// @Failure       400               {object}    dto.Error
// @Failure       404               {object}    dto.Error
// @Failure       500               {object}    dto.Error
// @Router        /products/{id}    [delete]
// @Security      ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := errors.New("ID inválido! Deve ser diferente de vazio.")
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if err = h.ProductDB.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}
