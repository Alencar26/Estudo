package services

import (
	"CRUD-Web/internal/entities"
	"CRUD-Web/internal/repositories"
	"CRUD-Web/pkg/httpcoffees"
	"encoding/json"
	"log"
)

type CoffeeService struct {
	Repo repositories.CoffeeRepository
}

// "construtor"
func NewCoffeeService(repo repositories.CoffeeRepository) *CoffeeService {
	return &CoffeeService{Repo: repo}
}

func (service *CoffeeService) GetAllCoffees() ([]entities.Coffee, error) {

	coffees, err := service.Repo.FindAll()
	if err != nil {
		log.Printf("Erro ao obter informações do repositório: %v\n", err)
		return nil, err
	}

	if len(coffees) < 1 {
		coffees, err = getCoffeesAPI()
		if err != nil {
			log.Printf("Erro ao chamar função getCoffeesAPI(): %v\n", err)
			return nil, err
		}
		return coffees, nil
	}

	return coffees, nil
}

func getCoffeesAPI() ([]entities.Coffee, error) {

	coffees := []entities.Coffee{}

	body, err := httpcoffees.GetCoffees()
	if err != nil {
		log.Printf("Erro ao chamar função httpcoffees.GetCoffees(): %v\n", err)
		panic(err.Error())
	}

	if err := json.Unmarshal(body, &coffees); err != nil {
		log.Printf("Erro no unmarshal do JSON para struct Coffee: %v\n", err)
		return nil, err
	}

	return coffees, nil
}

func (service *CoffeeService) CreateCoffee(coffee *entities.Coffee) error {
	if err := service.Repo.Create(coffee); err != nil {
		log.Printf("Erro ao criar um Coffee na base de dados: %v\n", err)
		return err
	}
	return nil
}
