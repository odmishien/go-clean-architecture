package registry

import (
	"haiken/entity"
	"haiken/infrastructure/logger"
	"haiken/infrastructure/onmemory"
	"haiken/interface/http/presenter"
	"haiken/usecase"
)

func NewUserInteractor() usecase.UserInteractor {
	return usecase.NewUserInteractor(presenter.NewUserPresenter(logger.NewLogger()), newUserService())
}

func newUserService() usecase.UserService {
	return usecase.NewUserService(onmemory.NewUserRepository(entity.Users{}))
}
