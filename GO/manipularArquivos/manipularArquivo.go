package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//CRIAR
	file, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//ESCREVER
	//quando eu sei que vou escrever uma string
	//tamanho, err := file.WriteString("Minha string.")
	//quando fou escrever qlqr coisa (binário)
	tamanho, err := file.Write([]byte("Escrevendo bytes..."))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Meu arquivo foi escrito e tem o tamanho de %d bytes.\n", tamanho)

	//FECHANDO ARQUIVO
	file.Close()

	//LER
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//LEITURA PARCIAL DO ARQUIVO PARA POUPAR MEMÓRIA EM ARQUIVOS GRANDES
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//pacote GO para ler arquivos em buffers
	reader := bufio.NewReader(arquivo2)
	//de quanto em quanto ele vai ler o arquivo?
	buffer := make([]byte, 10) //ler de 10 em 10 bytes

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	//REMOVER ARQUIVO
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
