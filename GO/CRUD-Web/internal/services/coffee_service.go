package services

import (
	"CRUD-Web/internal/entities"
	"CRUD-Web/internal/repositories"
	"CRUD-Web/pkg/httpcoffees"
	"encoding/json"
	"log"
)

type CoffeeService struct {
	repo repositories.CoffeeRepository
}

// "construtor"
func NewCoffeeService(repo repositories.CoffeeRepository) *CoffeeService {
	return &CoffeeService{repo: repo}
}

func (service *CoffeeService) GetAllCoffees() ([]entities.Coffee, error) {

	coffees, err := service.repo.FindAll()
	if err != nil {
		log.Printf("Erro ao obter informações do repositório: %V\n", err)
		return nil, err
	}

	if len(coffees) < 1 {
		coffees, err = getCoffeesAPI()
		if err != nil {
			log.Printf("Erro ao chamar função getCoffeesAPI(): %V\n", err)
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
		log.Printf("Erro ao chamar função httpcoffees.GetCoffees(): %V\n", err)
		panic(err.Error())
	}

	if err := json.Unmarshal(body, &coffees); err != nil {
		log.Printf("Erro no unmarshal do JSON para struct Coffee: %V\n", err)
		return nil, err
	}

	return coffees, nil
}

func (service *CoffeeService) CreateCoffee(coffee *entities.Coffee) error {
	if err := service.repo.Create(coffee); err != nil {
		log.Printf("Erro ao criar um Coffee na base de dados: %V\n", err)
		return err
	}
	return nil
}

func (service *CoffeeService) DeleteCoffee(id int) error {
	if err := service.repo.Delete(id); err != nil {
		log.Printf("Erro ao efetuar deleção do café com id %d: %V\n", id, err)
		return err
	}
	return nil
}

func (service *CoffeeService) GetCoffeeById(id int) (*entities.Coffee, error) {
	coffee, err := service.repo.FindById(id)
	if err != nil {
		log.Printf("Erro ao obter um caté pelo ID: %V", err)
		return nil, err
	}
	return coffee, nil
}

func (service *CoffeeService) UpdateCoffee(coffee *entities.Coffee) error {
	if err := service.repo.Update(coffee); err != nil {
		log.Printf("Erro ao atualizar um Coffee na base de dados: %V\n", err)
		return err
	}
	return nil
}
