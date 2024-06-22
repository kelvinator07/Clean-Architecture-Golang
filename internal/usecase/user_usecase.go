package usecase

import (
	"github.com/kelvinator05/clean-architecture-go/internal/entity"
	_interface "github.com/kelvinator05/clean-architecture-go/internal/interface"
)

type UserUseCase struct {
	Repo _interface.UserRepository
}

func (u *UserUseCase) CreateUser(name, email string) (*entity.User, error) {
	user := entity.NewUser(name, email)
	return u.Repo.Save(user)
}

func (u *UserUseCase) GetUserByID(id int) (*entity.User, error) {
	return u.Repo.FindByID(id)
}

func (u *UserUseCase) GetUserByEmail(email string) (*entity.User, error) {
	return u.Repo.FindByEmail(email)
}

func (u *UserUseCase) GetAllUsers() ([]*entity.User, error) {
	return u.Repo.GetAll()
}
