package usecase

import (
	"context"
	"haiken/usecase/output"
)

type UserPresenter interface {
	ViewAll(ctx context.Context, output *output.UserGetAllOutputData)
	ViewError(ctx context.Context, err error)
}
