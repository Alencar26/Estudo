package main

func main() {

	//já definindo o tipo
	const x int = 64

	//deixando com tipo indefinido. (o programa só vai atribuir um tipo quando for usar a constante.)
	const y = "Minha constante."

	//OBS: O tipo de uma constante só é definido no momento do uso
	//OBS: O tipo de uma variável é difinida no momento da atribuição/declaração.

	//várias constantes:
	const (
		a = 12
		b = "13"
		c = true
		d = false
		e = 14.1
	)

}
