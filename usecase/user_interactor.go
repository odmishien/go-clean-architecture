package usecase

import (
	"context"
	"haiken/usecase/input"
	"haiken/usecase/output"
)

type UserInteractor interface {
	GetAll(ctx context.Context)
	Create(ctx context.Context, input input.UserCreateInputData)
	GetByID(input input.UserGetByIDInputData)
}

type UserInteractorImpl struct {
	presenter UserPresenter
	service   UserService
}

func NewUserInteractor(presenter UserPresenter, service UserService) UserInteractor {
	return &UserInteractorImpl{
		presenter: presenter,
		service:   service,
	}
}

func (i *UserInteractorImpl) GetAll(ctx context.Context) {
	users, err := i.service.GetAll(ctx)
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
	createdID, err := i.service.Create(ctx, input)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}
	o := &output.UserCreateOutputData{CreatedID: createdID}
	i.presenter.ViewCreate(ctx, o)
}

func (i *UserInteractorImpl) GetByID(input input.UserGetByIDInputData) {
}
