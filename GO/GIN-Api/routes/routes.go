package routes

import (
	"GIN-Api/handles"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", handles.GetAllAlunos)
	r.GET("/:nome", handles.GetSaudacao)
	r.POST("/alunos", handles.CreateAluno)
	r.GET("/alunos/:id", handles.GetAluno)
	r.DELETE("/alunos/:id", handles.DeleteAluno)
	r.PUT("/alunos/:id", handles.UpdateAluno)

	r.Run(":5000")
}
