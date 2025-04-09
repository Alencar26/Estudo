package main

import (
	"GIN-Api/database"
	"GIN-Api/handles"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutesTest() *gin.Engine {
	r := gin.Default()

	return r
}

func TestVerificaSaudacaoStatusCode(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:nome", handles.GetSaudacao)
	req, _ := http.NewRequest(http.MethodGet, "/Nome de Teste", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "StatusCode deve ser 200")
}

func TestVerificaSaudacaoBody(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:nome", handles.GetSaudacao)
	req, _ := http.NewRequest(http.MethodGet, "/Nome de Teste", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	mockResponse := `{"code":200,"message":"Hello Nome de Teste","status":"success"}`

	assert.Equal(t, http.StatusOK, response.Code, "StatusCode deve ser 200")
	assert.Equal(t, mockResponse, response.Body.String(), "Body deve ser: "+mockResponse)
}

func TestGetAllAlunos(t *testing.T) {
	database.ConnectDB()
	r := SetupRoutesTest()
	r.GET("/alunos", handles.GetAllAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotEqual(t, "[]", response.Body.String(), "Retorno não deve ser vazio.")
}
