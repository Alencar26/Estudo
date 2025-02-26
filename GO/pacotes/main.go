package main

//usando pacotes externos
//comando: go mod init github.com/alencar26/convenção_da_comunidade
//por convenção usa-se um link do github
//comando para atualizar o go mod: go mod tidy
import (
	"fmt"

	"EstudoGO/pacotes/calculo"
)

func main() {
	s := calculo.Soma(12, 5)
	fmt.Printf("Resultado da soma: %v", s)
}
