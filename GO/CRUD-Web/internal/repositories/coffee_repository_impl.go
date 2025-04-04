package repositories

import (
	"CRUD-Web/internal/entities"
	"database/sql"
	"encoding/json"
	"log"
)

type CoffeeRepositoryImpl struct {
	db *sql.DB
}

// injeção de dependência - ""construtor""
func NewCoffeeRepository(db *sql.DB) *CoffeeRepositoryImpl {
	return &CoffeeRepositoryImpl{db: db}
}

func (r *CoffeeRepositoryImpl) Create(coffee *entities.Coffee) error {

	query := "INSERT INTO coffee(title, description, price, ingredients, image) VALUES($1, $2, $3, $4, $5)"
	insert, err := r.db.Prepare(query)
	if err != nil {
		log.Printf("Erro ao preparar query para insert: %V\n", err)
		return err
	}

	jsonIngredients, err := json.Marshal(coffee.Ingredients)
	if err != nil {
		log.Printf("Erro ao converter slice de String em JSON: %V\n", err)
		return err
	}

	result, err := insert.Exec(
		coffee.Title,
		coffee.Description,
		coffee.Price,
		jsonIngredients,
		coffee.Image,
	)
	defer r.db.Close()

	if err != nil {
		log.Printf("Erro ao executar query de insert: %V\n", err)
		return err
	} else {
		log.Printf("Insert realizado com sucesso: %V\n", result)
	}

	return nil
}

func (r *CoffeeRepositoryImpl) FindById(id int) (*entities.Coffee, error) {

	coffee := entities.Coffee{}
	ingredientsJSON := []byte{} //json

	query := "SELECT * FROM coffee WHERE id = $1;"
	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&coffee.Id,
		&coffee.Title,
		&coffee.Description,
		&coffee.Price,
		&ingredientsJSON,
		&coffee.Image,
	)
	if err != nil {
		log.Printf("Erro fazer scan para obter um café pelo ID: %V", err)
		return nil, err
	}

	err = json.Unmarshal(ingredientsJSON, &coffee.Ingredients)
	if err != nil {
		log.Printf("Falha no unmarshal dos ingredientes do banco de dados (JSON) para struct  coffee.Ingredients: %V\n", err)
		return nil, err
	}

	return &coffee, nil
}

func (r *CoffeeRepositoryImpl) FindAll() ([]entities.Coffee, error) {

	query := "SELECT * FROM coffee;"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Falha ao efetuar query SELECT no banco de dados: %V\n", err)
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
			log.Printf("Falha ao ler registros no banco de dados: %V\n", err)
			return nil, err
		}

		err = json.Unmarshal(ingredientsJSON, &coffee.Ingredients)
		if err != nil {
			log.Printf("Falha no unmarshal dos ingredientes do banco de dados (JSON) para struct  coffee.Ingredients: %V\n", err)
			return nil, err
		}

		coffees = append(coffees, coffee)
	}

	return coffees, nil
}

func (r *CoffeeRepositoryImpl) Update(coffee *entities.Coffee) error {

	query := "UPDATE coffee SET title=$1, description=$2, price=$3 WHERE id = $4;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Printf("Erro ao preparar query para update: %V\n", err)
		return err
	}

	result, err := stmt.Exec(coffee.Title, coffee.Description, coffee.Price, coffee.Id)
	defer r.db.Close()

	if err != nil {
		log.Printf("Erro ao executar query de update: %V\n", err)
		return err
	} else {
		log.Printf("Update realizado com sucesso: %V\n", result)
	}
	return nil
}

func (r *CoffeeRepositoryImpl) Delete(id int) error {

	query := "DELETE FROM coffee WHERE id = $1"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Printf("Erro ao preparar query para delete: %V\n", err)
		return err
	}

	result, err := stmt.Exec(id)
	defer r.db.Close()

	if err != nil {
		log.Printf("Erro ao executar query de delete: %V\n", err)
		return err
	} else {
		log.Printf("Delete realizado com sucesso: %V\n", result)
	}

	return nil
}
