package registry

import (
	"haiken/infrastructure/onmemory"
	"haiken/usecase"
)

func NewUserInteractor() usecase.UserInteractor {
	return usecase.NewUserInteractor(onmemory.NewUserRepository())
}
