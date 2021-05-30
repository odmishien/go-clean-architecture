package onmemory

import (
	"context"
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

func (r *UserRepositoryImpl) GetByID(id int) (*entity.User, error) {
	return nil, nil
}
