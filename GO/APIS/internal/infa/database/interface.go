package database

import "github.com/Alencar26/Estudo/GO/APIS/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
