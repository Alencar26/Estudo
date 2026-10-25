package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Alencar26/Estudo/GO/SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {

	ctx := context.Background() //contexto em branco
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Teste SQLC",
		Description: sql.NullString{String: "Desc Teste SQLC", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.ID, c.Name, c.Description)
	}
}
