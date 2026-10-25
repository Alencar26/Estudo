package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//USANDO  fmt.Scan
	fmt.Println("fmt.Scan ----------------------------")

	//Observação:  Só lê o input até o primeiro espaço.
	var nome string
	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)
	fmt.Printf("Seu nome é %s.\n", nome)

	//------------------------------------------------------

	//USANDO fmt.Scanln
	fmt.Println("\nfmt.Scanln ----------------------------")

	//Observação: Só lê até a quebra de linha
	var nome2 string
	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome2)
	fmt.Printf("Seu nome é %s.\n", nome2)

	//------------------------------------------------------

	//USANDO bufio.Reader
	fmt.Println("\nbufio.Reader ----------------------------")

	//Observação: MELHOR OPÇÃO. (Lê a linha inteira, incluindo espaços)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite uma frase:")
	texto, _ := reader.ReadString('\n') //informando que quero ler até a quebra de linha
	texto = strings.TrimSpace(texto)    //remove o espaço(quebra de linha) final
	fmt.Printf("Sua frase é: %s.\n", texto)
}
