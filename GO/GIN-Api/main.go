package main

import (
	"GIN-Api/database"
	"GIN-Api/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
