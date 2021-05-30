package usecase

import (
	"context"
	"haiken/entity"
	"haiken/usecase/input"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	ValidateCreateInput(input input.UserCreateInputData) error
	GetAll(ctx context.Context) (entity.Users, error)
	Create(ctx context.Context, input input.UserCreateInputData) (createdID int, err error)
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) ValidateCreateInput(input input.UserCreateInputData) error {
	validate := validator.New()
	return validate.Struct(input)
}

func (s *UserServiceImpl) GetAll(ctx context.Context) (entity.Users, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserServiceImpl) Create(ctx context.Context, input input.UserCreateInputData) (createdID int, err error) {
	return s.repo.Create(ctx, input.Name, input.Email, input.Password)
}
