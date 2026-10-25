package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         Saldo
}

func (c *ContaPoupanca) Transferir(valor Saldo, contaDestino *ContaPoupanca) {
	if valor > 0 && c.saldo.IsPositive() {
		c.saldo -= valor
		contaDestino.saldo += valor
	} else {
		fmt.Printf("Transferência de %.2f falhou\n", valor)
	}
}

func (c *ContaPoupanca) Depositar(valor Saldo) Saldo {
	if valor > 0 {
		c.saldo += valor
	} else {
		fmt.Printf("Depósito de %.2f falhou\n", valor)
	}
	return c.saldo
}

func (c *ContaPoupanca) Sacar(valor Saldo) Saldo {
	if c.saldo.IsPositive() && valor > 0 {
		c.saldo -= valor
	} else {
		fmt.Printf("saldo insuficiente - Seu saque de %.2f falhou\n", valor)
		fmt.Printf("saldo atual: %.2f\n", c.saldo)
	}
	return c.saldo
}

func (c *ContaPoupanca) SaldoEhPositivo() bool {
	return c.saldo.IsPositive()
}

func (c *ContaPoupanca) ConsultaSaldo() {
	fmt.Printf("saldo atual: %.2f\n", c.saldo)
}
