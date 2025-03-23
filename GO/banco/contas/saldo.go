package contas

type Saldo float64

func (s Saldo) IsPositive() bool {
	return s >= 0
}
