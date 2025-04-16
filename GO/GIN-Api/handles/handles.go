package handles

import (
	"GIN-Api/database"
	"GIN-Api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func GetSaudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + nome,
		"status":  "success",
		"code":    200,
	})
}

func CreateAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf("Erro ao criar aluno: %v\n", err)
		return
	}

	if err := aluno.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf("Erro na validação do Aluno: %v\n", err)
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Aluno criado com sucesso!",
		"aluno":   aluno,
		"status":  "success",
		"code":    201,
	})
}

func GetAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
			"status":  "error",
			"code":    404,
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	id := c.Params.ByName("id")

	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
			"status":  "error",
			"code":    404,
		})
		return
	}
	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno deletado com sucesso!",
		"status":  "success",
		"code":    200,
	})
}

func UpdateAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
			"status":  "error",
			"code":    404,
		})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("Erro ao atualizar aluno: %v\n", err)
		return
	}

	if err := aluno.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf("Erro na validação do Aluno: %v\n", err)
		return
	}

	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno atualizado com sucesso!",
		"status":  "success",
		"code":    200,
	})
}

func GetAlunoByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var aluno models.Aluno

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
			"status":  "error",
			"code":    404,
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaHTML(c *gin.Context) {

	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
