package usecase

import (
	"context"
)

type UserInteractor interface {
	GetAll(ctx context.Context)
	Create(input UserCreateInputData)
	GetByID(input UserGetByIDInputData)
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
		i.presenter.ViewAll(ctx, nil, err)
	}
	o := &UserGetAllOutputData{
		Users: users,
	}
	i.presenter.ViewAll(ctx, o, nil)
}
func (i *UserInteractorImpl) Create(input UserCreateInputData) {
}

func (i *UserInteractorImpl) GetByID(input UserGetByIDInputData) {
}
