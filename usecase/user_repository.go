package usecase

import "haiken/entity"

type UserRepository interface {
	GetAll() (entity.Users, error)
	Create(entity.User) (id int, err error)
	GetByID(id int) (*entity.User, error)
}
