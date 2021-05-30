package usecase

import (
	"context"
	"haiken/entity"
)

type UserRepository interface {
	GetAll(ctx context.Context) (entity.Users, error)
	Create(entity.User) (id int, err error)
	GetByID(id int) (*entity.User, error)
}
