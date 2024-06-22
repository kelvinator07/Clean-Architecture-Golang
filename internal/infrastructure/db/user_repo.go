package db

import (
	"errors"
	"strconv"

	"github.com/kelvinator05/clean-architecture-go/internal/entity"
)

type InMemoryUserRepo struct {
	users []*entity.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: make([]*entity.User, 0)}
}

func (repo *InMemoryUserRepo) Save(user *entity.User) (*entity.User, error) {
	for _, u := range repo.users {
		if u.Email == user.Email {
			return nil, errors.New("user already exists with email " + user.Email)
		}
	}
	user.ID = len(repo.users) + 1
	repo.users = append(repo.users, user)
	return user, nil
}

func (repo *InMemoryUserRepo) FindByID(id int) (*entity.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found with id " + strconv.Itoa(id))
}

func (repo *InMemoryUserRepo) FindByEmail(email string) (*entity.User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found with email " + email)
}

func (repo *InMemoryUserRepo) GetAll() ([]*entity.User, error) {
	return repo.users, nil
}
