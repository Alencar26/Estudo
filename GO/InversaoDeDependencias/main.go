package main

type Order struct {
	Id    int
	Valor float64
}

type MercadoPago struct{}
type Stripe struct{}
type Pagarme struct{}

type GatewayPagamento interface {
	Pay(ordem *Order)
}

func (mp MercadoPago) Pay(ordem *Order) {
	//lógica
}

func (s Stripe) Pay(ordem *Order) {
	//lógica
}

func (p Pagarme) Pay(ordem *Order) {
	//lógica
}

func ProcessaPagamento(ordem Order, gatewayPagamentoInterface GatewayPagamento) {
	gatewayPagamentoInterface.Pay(&ordem)
}

func main() {

	ordem := Order{Id: 1, Valor: 2.5}

	//vários meios de pagamento
	stripe := Stripe{}
	mercadoPago := MercadoPago{}
	pagarme := Pagarme{}

	//posso usar qualquer um deles, pois todos implementam a mesma interface
	//posso trocar qual o gtw de pagamento sem complicações
	ProcessaPagamento(ordem, stripe)
	ProcessaPagamento(ordem, mercadoPago)
	ProcessaPagamento(ordem, pagarme)
}
