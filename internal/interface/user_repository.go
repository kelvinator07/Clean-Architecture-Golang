package _interface

import "github.com/kelvinator05/clean-architecture-go/internal/entity"

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindByID(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
}
