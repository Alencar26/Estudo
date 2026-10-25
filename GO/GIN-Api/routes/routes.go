package routes

import (
	"GIN-Api/handles"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")   //para conseguir renderizar as páginas HTML
	r.Static("/assets", "./assets") //para servir arquivos estáticos (css, js, imagens, etc.)

	r.GET("/alunos", handles.GetAllAlunos)
	r.GET("/:nome", handles.GetSaudacao)
	r.POST("/alunos", handles.CreateAluno)
	r.GET("/alunos/:id", handles.GetAluno)
	r.DELETE("/alunos/:id", handles.DeleteAluno)
	r.PUT("/alunos/:id", handles.UpdateAluno)
	r.GET("/alunos/cpf/:cpf", handles.GetAlunoByCPF)
	r.GET("/index", handles.ExibePaginaHTML) //para servir a página inicial
	r.NoRoute(handles.RotaNaoEncontrada404)

	r.Run(":5000")
}
