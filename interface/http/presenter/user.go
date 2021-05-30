package presenter

import (
	"context"
	"encoding/json"
	"haiken/usecase"
	"haiken/usecase/output"
	"log"
	"net/http"
)

type UserPresenterImpl struct{}

func NewUserPresenter() usecase.UserPresenter {
	return &UserPresenterImpl{}
}

func (p *UserPresenterImpl) ViewAll(ctx context.Context, output *output.UserGetAllOutputData) {
	w := getResponseWriter(ctx)
	returnJSON(w, output)
}

func (p *UserPresenterImpl) ViewCreate(ctx context.Context, output *output.UserCreateOutputData) {
	w := getResponseWriter(ctx)
	returnJSON(w, output)
}

func (p *UserPresenterImpl) ViewError(ctx context.Context, err error) {
	w := getResponseWriter(ctx)
	log.Fatalf("error: %s", err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func returnJSON(w http.ResponseWriter, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
