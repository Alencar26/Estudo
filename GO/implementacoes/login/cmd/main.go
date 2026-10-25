package main

import (
	"fmt"

	"github.com/Alencar26/Estudo/GO/implementacoes/login/internal/model"
	hashing "github.com/Alencar26/Estudo/GO/implementacoes/login/pkg/hash"
)

func main() {

	hash, err := hashing.HashPassword("senha123")
	if err != nil {
		return
	}

	user := model.NewUser("email@example.com", hash)

	fmt.Println(user)

	isValid := hashing.VerifyPassword("senha123", []byte(hash), nil)
	fmt.Println(isValid)
}
