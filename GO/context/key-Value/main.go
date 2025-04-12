package main

import (
	"context"
	"fmt"
)

func main() {

	//É possível passar metadados por contexto.
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "Bearer token12345")
	reservaHotel(ctx)
}

func reservaHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
