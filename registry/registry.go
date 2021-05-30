package registry

import (
	"haiken/infrastructure/onmemory"
	"haiken/interface/http/presenter"
	"haiken/usecase"
)

func NewUserInteractor() usecase.UserInteractor {
	return usecase.NewUserInteractor(onmemory.NewUserRepository(), presenter.NewUserPresenter())
}
