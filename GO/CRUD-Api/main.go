package main

import (
	"CRUD-Api/database"
	"CRUD-Api/routes"
	"fmt"
)

func main() {
	database.ConnectionDB()

	fmt.Println("Run Server...")
	routes.HandlerRequest()
}
