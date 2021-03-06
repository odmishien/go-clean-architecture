package usecase

import (
	"context"
	"haiken/usecase/input"
	"haiken/usecase/output"
)

type UserInteractor interface {
	GetAll(ctx context.Context)
	Create(ctx context.Context, input input.UserCreateInputData)
	GetByID(ctx context.Context, input input.UserGetByIDInputData)
}

type UserInteractorImpl struct {
	repo      UserRepository
	presenter UserPresenter
	service   UserService
}

func NewUserInteractor(repo UserRepository, presenter UserPresenter, service UserService) UserInteractor {
	return &UserInteractorImpl{
		repo:      repo,
		presenter: presenter,
		service:   service,
	}
}

func (i *UserInteractorImpl) GetAll(ctx context.Context) {
	users, err := i.repo.GetAll(ctx)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}
	o := &output.UserGetAllOutputData{
		Users: users,
	}
	i.presenter.ViewAll(ctx, o)
}
func (i *UserInteractorImpl) Create(ctx context.Context, input input.UserCreateInputData) {
	err := i.service.ValidateCreateInput(input)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}
	createdID, err := i.repo.Create(ctx, input.Name, input.Email, input.Password)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}
	o := &output.UserCreateOutputData{CreatedID: createdID}
	i.presenter.ViewCreate(ctx, o)
}

func (i *UserInteractorImpl) GetByID(ctx context.Context, input input.UserGetByIDInputData) {
	u, err := i.repo.GetByID(ctx, input.ID)
	if err != nil {
		i.presenter.ViewError(ctx, err)
	}
	o := &output.UserGetByIDOutputData{
		User: *u,
	}
	i.presenter.View(ctx, o)
}
