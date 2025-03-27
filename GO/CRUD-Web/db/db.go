package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// retorna a conexão com db
func ConnectDB() *sql.DB {
	// connection := "user dbname password host port sslmode"
	connection := "user=al3ncar dbname=estudo_db password=postgresSQL host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Printf("Falha na conexão com banco de dados: %v\n", err)
		panic(err.Error())
	}
	return db
}
