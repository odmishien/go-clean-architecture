package onmemory

import (
	"haiken/entity"
	"haiken/usecase"
)

type UserRepositoryImpl struct {
	db entity.Users
}

func NewUserRepository() usecase.UserRepository {
	return &UserRepositoryImpl{
		db: []entity.User{},
	}
}

func (r *UserRepositoryImpl) Create(entity.User) (id int, err error) {
	return 0, nil
}

func (r *UserRepositoryImpl) GetByID(id int) (*entity.User, error) {
	return nil, nil
}
