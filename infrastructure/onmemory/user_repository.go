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
		db: []entity.User{
			{
				ID:       1,
				Name:     "hoge",
				Email:    "hoge.gmail.com",
				Password: "asdfaygaerga",
			},
		},
	}
}

func (r *UserRepositoryImpl) GetAll() (entity.Users, error) {
	return r.db, nil
}

func (r *UserRepositoryImpl) Create(entity.User) (id int, err error) {
	return 0, nil
}

func (r *UserRepositoryImpl) GetByID(id int) (*entity.User, error) {
	return nil, nil
}
