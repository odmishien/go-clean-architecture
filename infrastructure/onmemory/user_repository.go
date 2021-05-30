package onmemory

import (
	"context"
	"errors"
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
				Name:     "default",
				Email:    "default@gmail.com",
				Password: "default",
			},
		},
	}
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) (entity.Users, error) {
	return r.db, nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, name, email, password string) (id int, err error) {
	cID := len(r.db)
	r.db = append(r.db, entity.User{
		ID:       cID + 1,
		Name:     name,
		Email:    email,
		Password: password,
	})
	return cID + 1, nil
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, id int) (*entity.User, error) {
	for _, u := range r.db {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
