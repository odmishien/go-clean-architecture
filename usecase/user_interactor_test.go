package usecase

import (
	"context"
	"errors"
	"haiken/entity"
	mock_usecase "haiken/mock"
	"haiken/usecase/output"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserInteractorImpl_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	us := entity.Users{
		{
			ID:   1,
			Name: "test1",
		},
		{
			ID:   2,
			Name: "test2",
		},
	}
	ctx := context.Background()
	tests := []struct {
		name                  string
		prepareMockRepository func(m *mock_usecase.MockUserRepository)
		prepareMockPresenter  func(m *mock_usecase.MockUserPresenter)
		args                  args
	}{
		{
			name: "OK",
			prepareMockRepository: func(m *mock_usecase.MockUserRepository) {
				m.EXPECT().GetAll(ctx).Return(us, nil)
			},
			prepareMockPresenter: func(m *mock_usecase.MockUserPresenter) {
				m.EXPECT().ViewAll(ctx, &output.UserGetAllOutputData{Users: us})
			},
			args: args{ctx: ctx},
		},
		{
			name: "something wrong",
			prepareMockRepository: func(m *mock_usecase.MockUserRepository) {
				m.EXPECT().GetAll(ctx).Return(nil, errors.New("something wrong"))
			},
			prepareMockPresenter: func(m *mock_usecase.MockUserPresenter) {
				m.EXPECT().ViewError(ctx, errors.New("something wrong"))
			},
			args: args{ctx: ctx},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepository := mock_usecase.NewMockUserRepository(mockCtrl)
			mockPresenter := mock_usecase.NewMockUserPresenter(mockCtrl)
			tt.prepareMockRepository(mockRepository)
			tt.prepareMockPresenter(mockPresenter)

			i := &UserInteractorImpl{
				repo:      mockRepository,
				presenter: mockPresenter,
				service:   NewUserService(),
			}
			i.GetAll(tt.args.ctx)
		})
	}
}
