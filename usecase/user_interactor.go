package usecase

import (
	"context"
	"haiken/usecase/input"
	"haiken/usecase/output"
)

type UserInteractor interface {
	GetAll(ctx context.Context)
	Create(input input.UserCreateInputData)
	GetByID(input input.UserGetByIDInputData)
}

type UserInteractorImpl struct {
	repo      UserRepository
	presenter UserPresenter
}

func NewUserInteractor(repo UserRepository, presenter UserPresenter) UserInteractor {
	return &UserInteractorImpl{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *UserInteractorImpl) GetAll(ctx context.Context) {
	users, err := i.repo.GetAll(ctx)
	if err != nil {
		i.presenter.ViewError(ctx, err)
	}
	o := &output.UserGetAllOutputData{
		Users: users,
	}
	i.presenter.ViewAll(ctx, o)
}
func (i *UserInteractorImpl) Create(input input.UserCreateInputData) {
}

func (i *UserInteractorImpl) GetByID(input input.UserGetByIDInputData) {
}
