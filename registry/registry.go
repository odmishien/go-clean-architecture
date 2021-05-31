package registry

import (
	"haiken/entity"
	"haiken/infrastructure/logger"
	"haiken/interface/gateway/onmemory"
	"haiken/interface/http/presenter"
	"haiken/usecase"
)

func NewUserInteractor() usecase.UserInteractor {
	return usecase.NewUserInteractor(
		onmemory.NewUserRepository(entity.Users{}),
		presenter.NewUserPresenter(logger.NewLogger()),
		usecase.NewUserService(),
	)
}
