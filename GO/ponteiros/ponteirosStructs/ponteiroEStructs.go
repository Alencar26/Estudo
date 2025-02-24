package main

type Conta struct {
	saldo int
}

// sem uso de ponteiro
func (c Conta) simularEmprestimo(valor int) int {
	c.saldo += valor
	return c.saldo
}

// uso de ponteiro
func (c *Conta) emprestar(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	println("Simulando: ", conta.simularEmprestimo(200))
	println("Saldo Atual", conta.saldo)

	println("")

	println("Emprestimo: ", conta.emprestar(200))
	println("Saldo Atual", conta.saldo)
}
