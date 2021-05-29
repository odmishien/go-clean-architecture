package usecase

import "haiken/entity"

type UserRepository interface {
	Create(entity.User) (id int, err error)
	GetByID(id int) (*entity.User, error)
}
