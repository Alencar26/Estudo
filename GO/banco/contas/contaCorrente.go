package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         Saldo
}

func (c *ContaCorrente) Transferir(valor Saldo, contaDestino *ContaCorrente) {
	if valor > 0 && c.saldo.IsPositive() {
		c.saldo -= valor
		contaDestino.saldo += valor
	} else {
		fmt.Printf("Transferência de %.2f falhou\n", valor)
	}
}

func (c *ContaCorrente) Depositar(valor Saldo) Saldo {
	if valor > 0 {
		c.saldo += valor
	} else {
		fmt.Printf("Depósito de %.2f falhou\n", valor)
	}
	return c.saldo
}

func (c *ContaCorrente) Sacar(valor Saldo) Saldo {
	if c.saldo.IsPositive() && valor > 0 {
		c.saldo -= valor
	} else {
		fmt.Printf("saldo insuficiente - Seu saque de %.2f falhou\n", valor)
		fmt.Printf("saldo atual: %.2f\n", c.saldo)
	}
	return c.saldo
}

func (c *ContaCorrente) SaldoEhPositivo() bool {
	return c.saldo.IsPositive()
}

func (c *ContaCorrente) ConsultaSaldo() {
	fmt.Printf("saldo atual: %.2f\n", c.saldo)
}
