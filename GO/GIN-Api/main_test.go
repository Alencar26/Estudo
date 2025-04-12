package main

import (
	"GIN-Api/database"
	"GIN-Api/handles"
	"GIN-Api/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	return r
}

var idAluno int

func criarAlunoMock() {
	aluno := models.Aluno{
		Nome: "Aluno Teste Automatizado",
		CPF:  "12345678900",
		RG:   "123456789",
	}
	database.DB.Create(&aluno)
	idAluno = int(aluno.ID)
}

func deletaAlunoMock() {
	database.DB.Delete(&models.Aluno{}, idAluno)
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
	criarAlunoMock()
	defer deletaAlunoMock()
	r := SetupRoutesTest()
	r.GET("/alunos", handles.GetAllAlunos)

	req, _ := http.NewRequest(http.MethodGet, "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotEqual(t, "[]", response.Body.String(), "Retorno não deve ser vazio.")
}

func TestGetAlunoById(t *testing.T) {
	database.ConnectDB()
	criarAlunoMock()
	defer deletaAlunoMock()

	r := SetupRoutesTest()
	r.GET("/alunos/:id", handles.GetAluno)

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/alunos/%d", idAluno), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var aluno models.Aluno
	if err := json.NewDecoder(response.Body).Decode(&aluno); err != nil {
		t.Errorf("Erro ao decodificar resposta JSON: %v", err)
	}

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, idAluno, int(aluno.ID), "Id deve ser igual ao id do aluno criado")
}

func TestUpdateAluno(t *testing.T) {
	database.ConnectDB()
	criarAlunoMock()
	defer deletaAlunoMock()

	r := SetupRoutesTest()
	r.PUT("/alunos/:id", handles.UpdateAluno)

	aluno := models.Aluno{
		Nome: "Caio Castro",
		CPF:  "00987654321",
		RG:   "987654321",
	}

	alunoJSON, err := json.Marshal(aluno)
	if err != nil {
		t.Errorf("Error: Falhou ao converter struct aluno para JSON: %v", err)
	}

	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/alunos/%d", idAluno), bytes.NewBuffer(alunoJSON))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	//terminar
	assert.Equal(t, int(aluno.ID), response.Body, msgAndArgs ...interface{})
}
