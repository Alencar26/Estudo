package main

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// altomaticamente qualquer struct que possuir a função "Desativar" estará implementando essa interface.
// na interface só podemos passar métodos
type Pessoa interface {
	Desativar()
}

// FUNÇÃO DA STRUCT "Método de Classe"
// * indica que estou usando um ponteiro para instância atual
func (c *Cliente) Desativar() {
	c.Ativo = false
	println(c.Ativo)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cliente1 := Cliente{
		Nome:  "André",
		Idade: 28,
		Ativo: true,
	}

	//cliente1 é do tipo "Pessoa" posi ele implementar o "método" Desativar
	Desativacao(&cliente1)
}
