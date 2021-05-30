package usecase

import "context"

type UserPresenter interface {
	ViewAll(ctx context.Context, output *UserGetAllOutputData, err error)
}
