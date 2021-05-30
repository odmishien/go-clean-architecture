package usecase

import (
	"context"
	"haiken/entity"
)

type UserRepository interface {
	GetAll(ctx context.Context) (entity.Users, error)
	Create(ctx context.Context, name, email, password string) (id int, err error)
	GetByID(ctx context.Context, id int) (*entity.User, error)
}
