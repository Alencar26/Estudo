package main

//usando pacotes externos
//comando: go mod init github.com/alencar26/convenção_da_comunidade
//por convenção usa-se um link do github
//comando para atualizar o go mod: go mod tidy (remover ou add pacotes)
import (
	"fmt"

	"EstudoGO/pacotes/calculo"

	//para instalar um pacote externo basta usar o comando abaixo:
	//comando: go get github.com/google/uuid
	"github.com/google/uuid"
)

func main() {
	s := calculo.Soma(12, 5)
	fmt.Printf("Resultado da soma: %v\n", s)

	fmt.Println("UUID:", uuid.New())
}
