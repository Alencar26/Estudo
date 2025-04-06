package database

import (
	"GIN-Api/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	connection := getStringConnectionDB()
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v\n", err)
	} else {
		log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	}

	DB.AutoMigrate(&models.Aluno{})
}

func getStringConnectionDB() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Falha ao recuperar variáveis de ambiente: %V\n", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUserName,
		dbPassword,
		dbDatabase,
		dbPort,
	)
}
