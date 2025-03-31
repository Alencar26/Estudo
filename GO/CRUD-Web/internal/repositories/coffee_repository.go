package repositories

import "CRUD-Web/internal/entities"

type CoffeeRepository interface {
	Create(coffee *entities.Coffee) error
	FindById(id int) (*entities.Coffee, error)
	FindAll() ([]entities.Coffee, error)
	Update(coffee *entities.Coffee) error
	Delete(id int) error
}
