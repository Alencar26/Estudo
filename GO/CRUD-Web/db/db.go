package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// retorna a conexão com db
func ConnectDB() *sql.DB {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Falha ao recuperar variáveis de ambiente: %V\n", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	// connection := "user dbname password host port sslmode"
	connection := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUserName,
		dbDatabase,
		dbPassword,
		dbHost,
		dbPort,
	)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Printf("Falha na conexão com banco de dados: %V\n", err)
		panic(err.Error())
	}
	return db
}
