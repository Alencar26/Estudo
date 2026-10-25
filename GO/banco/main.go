package main

import (
	"banco/clientes"
	c "banco/contas"
	"fmt"
)

func PagarBoleto(conta c.VerificarConta, valor c.Saldo) {
	conta.Sacar(valor)
}

func main() {
	conta := c.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "João",
			Profissao: "Desenvolvedor",
			CPF:       "560.296.770-20",
		},
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}

	conta2 := c.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Maria",
			Profissao: "Desenvolvedora",
			CPF:       "123.2A2.739-14",
		},
		NumeroAgencia: 589,
		NumeroConta:   123457,
	}

	contaPoup := c.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "José",
			Profissao: "Desenvolvedora",
			CPF:       "123.232.739-15",
		},
		NumeroAgencia: 589,
		NumeroConta:   123457,
	}

	conta.Depositar(-100)
	conta.Depositar(100)
	conta2.Depositar(500)

	fmt.Println("Valida CPF:", conta.Titular.CPF.ValidaCPF())
	fmt.Println("Valida CPF:", conta2.Titular.CPF.ValidaCPF())

	fmt.Println("--------------- MANIPULAÇÕES -----------------")

	fmt.Println(conta.SaldoEhPositivo())
	fmt.Println(conta.Depositar(200))
	fmt.Println(conta.SaldoEhPositivo())
	fmt.Println(conta.Sacar(50))
	fmt.Println(conta.Sacar(251))
	fmt.Println(conta.SaldoEhPositivo())
	fmt.Println(conta.Sacar(250))
	fmt.Println(conta.SaldoEhPositivo())
	fmt.Println(conta.Sacar(-1000))
	fmt.Println(conta.SaldoEhPositivo())
	fmt.Println(conta.Depositar(-2000))
	fmt.Println(conta.SaldoEhPositivo())
	fmt.Println(conta.Depositar(2000))
	fmt.Println(conta.SaldoEhPositivo())
	conta.ConsultaSaldo()
	conta2.ConsultaSaldo()
	conta.Transferir(1000, &conta2)
	conta.ConsultaSaldo()
	conta2.ConsultaSaldo()

	fmt.Println("--------------- PAGAMENTO BOLETO -----------------")

	fmt.Println("Conta Corrente")

	conta.ConsultaSaldo()
	PagarBoleto(&conta, 100)
	conta.ConsultaSaldo()

	fmt.Println("Conta Poupança")
	contaPoup.Depositar(200)
	contaPoup.ConsultaSaldo()
	PagarBoleto(&contaPoup, 100)
	contaPoup.ConsultaSaldo()
}
