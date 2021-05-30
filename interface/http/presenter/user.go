package presenter

import (
	"context"
	"encoding/json"
	"haiken/entity"
	"haiken/usecase"
	"haiken/usecase/output"
	"net/http"
)

type UserPresenterImpl struct {
	logger entity.Logger
}

func NewUserPresenter(logger entity.Logger) usecase.UserPresenter {
	return &UserPresenterImpl{
		logger: logger,
	}
}

func (p *UserPresenterImpl) ViewAll(ctx context.Context, output *output.UserGetAllOutputData) {
	w := getResponseWriter(ctx)
	p.returnJSON(w, output)
}

func (p *UserPresenterImpl) View(ctx context.Context, output *output.UserGetByIDOutputData) {
	w := getResponseWriter(ctx)
	p.returnJSON(w, output)
}

func (p *UserPresenterImpl) ViewCreate(ctx context.Context, output *output.UserCreateOutputData) {
	w := getResponseWriter(ctx)
	p.returnJSON(w, output)
}

func (p *UserPresenterImpl) ViewError(ctx context.Context, err error) {
	w := getResponseWriter(ctx)
	p.logger.Errorf(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (p *UserPresenterImpl) returnJSON(w http.ResponseWriter, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		p.logger.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
