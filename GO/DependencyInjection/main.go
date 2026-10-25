package main

import (
	"database/sql"

	"github.com/Alencar26/Estudo/GO/DependencyInjection/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		panic(err)
	}

	repository := product.NewProductRepository(db)

	useCase := product.NewProductUseCase(repository)

	product, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	println(product.Name)
}
