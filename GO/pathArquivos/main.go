package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//obter o path qye meu binário está rodando
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	fmt.Println(path)

	//pega somente o path diretório dele
	dir := filepath.Dir(path)

	fmt.Println(dir)

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	fmt.Println("ARQUIVOS DESSE DIRETÓRIO:")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
