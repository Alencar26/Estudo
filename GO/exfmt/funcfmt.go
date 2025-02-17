package main

import "fmt"

func main() {
	//------------------------------------------

	//PACOTE FMT

	//string interpretada: \n \t \"  não vão aparecer na string.
	s := "Oi, bom dia\nComo você vai?\t \"espero que tudo bem\"."
	fmt.Println(s)

	//raw string literals: é uma string que vai mostrar tudo, até caracterar especiais que seriam interpretados
	r := `"Oi, bom dia\nComo
  lalala
              você vai?\t \"es
    pero que tudo bem\"."`
	fmt.Println(r)

	//cada caracter de uma string é chamada de rune.

	fmt.Print("Sem quebra de linha")
	fmt.Println("Com quebra de linha")
	fmt.Printf("Podemos inserir formatação no texto\n")

	f := "oi"
	g := "como vai?"

	//Valor é salvo em uma string ao invés de ser jogado em tela.
	h := fmt.Sprint("Concatenar", " ", "strings", " diferentes", " variáveis tbm: ", f, g, "?\n")
	i := fmt.Sprintf("variação")
	j := fmt.Sprintln("Variação")
	fmt.Print(h)
	fmt.Print(i)
	fmt.Print(j)

	//Para escrever em arquivos F = file
	//fmt.Fprint()
	//fmt.Fprintf()
	//fmt.Fprintln()

	//-----------------------------------
	//------------------------------------------

	//PACOTE FMT

	//string interpretada: \n \t \"  não vão aparecer na string.
	s = "Oi, bom dia\nComo você vai?\t \"espero que tudo bem\"."
	fmt.Println(s)

	//raw string literals: é uma string que vai mostrar tudo, até caracterar especiais que seriam interpretados
	r = `"Oi, bom dia\nComo
  lalala
              você vai?\t \"es
    pero que tudo bem\"."`
	fmt.Println(r)

	//cada caracter de uma string é chamada de rune.

	fmt.Print("Sem quebra de linha")
	fmt.Println("Com quebra de linha")
	fmt.Printf("Podemos inserir formatação no texto\n")

	f = "oi"
	g = "como vai?"

	//Valor é salvo em uma string ao invés de ser jogado em tela.
	h = fmt.Sprint("Concatenar", " ", "strings", " diferentes", " variáveis tbm: ", f, g, "?\n")
	i = fmt.Sprintf("variação")
	j = fmt.Sprintln("Variação")
	fmt.Print(h)
	fmt.Print(i)
	fmt.Print(j)

	//Para escrever em arquivos F = file
	//fmt.Fprint()
	//fmt.Fprintf()
	//fmt.Fprintln()

	//-----------------------------------

}
