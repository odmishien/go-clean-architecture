package usecase

import (
	"haiken/usecase/input"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	ValidateCreateInput(input input.UserCreateInputData) error
}

type UserServiceImpl struct{}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) ValidateCreateInput(input input.UserCreateInputData) error {
	validate := validator.New()
	return validate.Struct(input)
}
