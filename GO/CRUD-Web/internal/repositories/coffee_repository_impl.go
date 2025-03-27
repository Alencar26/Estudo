package repositories

import (
	"CRUD-Web/internal/entities"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type CoffeeRepositoryImpl struct {
	db *sql.DB
}

// injeção de dependência
func NewCoffeeRepository(db *sql.DB) *CoffeeRepositoryImpl {
	return &CoffeeRepositoryImpl{db: db}
}

func (r *CoffeeRepositoryImpl) Create(coffee *entities.Coffee) error {
	//TODO
	return nil
}

func (r *CoffeeRepositoryImpl) FindById(id int) (*entities.Coffee, error) {
	//TODO
	return nil, nil
}
func (r *CoffeeRepositoryImpl) FindAll() ([]entities.Coffee, error) {

	query := "SELECT * FROM coffee;"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Falha ao efetuar query SELECT no banco de dados: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	coffees := []entities.Coffee{}

	if !rows.Next() {
		return coffees, nil
	}

	for rows.Next() {
		coffee := entities.Coffee{}
		ingredientsJSON := []byte{}

		err = rows.Scan(
			&coffee.Id,
			&coffee.Title,
			&coffee.Description,
			&coffee.Price,
			&ingredientsJSON,
			&coffee.Image,
		)
		if err != nil {
			log.Printf("Falha ao ler registros no banco de dados: %v\n", err)
			return nil, err
		}

		err = json.Unmarshal(ingredientsJSON, &coffee.Ingredients)
		if err != nil {
			log.Printf("Falha no unmarshal dos ingredientes do banco de dados (JSON) para struct  coffee.Ingredients: %v\n", err)
			return nil, err
		}

		coffees = append(coffees, coffee)
		fmt.Println(coffee)
	}

	return coffees, nil
}

func (r *CoffeeRepositoryImpl) Delete(id int) error {
	//TODO
	return nil
}
