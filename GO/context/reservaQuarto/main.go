package main

import (
	"context"
	"fmt"
	"time"
)

//Contexto serve para controlar o fluxo de execução de um programa em Go.
//É uma estrutura que permite cancelar operações assíncronas
//E compartilhar informações entre goroutines.

//Opção de um contexto:
//WithCancel - Cancela o contexto quando cancel é chamado
//WithDeadline - Cancela o contexto quando o deadline é atingido
//WithTimeout - Cancela o contexto após o timeout
//WithValue - Quando passa o valor para uma chave

func main() {

	ctx := context.Background()                            //inicializando contexto
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) //Adicionando regra de expiração
	defer cancel()                                         //cancelae contexto no final
	reservaQuarto(ctx)                                     //chamei minha função passando contexto
}

func reservaQuarto(ctx context.Context) {
	select {
	case <-ctx.Done(): //Se passar o timeout de 3 segundos cai nessa condição (quando cxt é cancelado ele recebe um Done)
		fmt.Println("Reserva cancelada. Timeout reached")
		return
	case <-time.After(1 * time.Second): //Se o tempo for de 1 segundo
		fmt.Println("Reserva confirmada.")
		return
	}
}
