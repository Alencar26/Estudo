package contas

type VerificarConta interface {
	Sacar(valor Saldo) Saldo
}
