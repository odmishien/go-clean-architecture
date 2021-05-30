package presenter

import (
	"context"
	"encoding/json"
	"haiken/usecase"
	"haiken/usecase/output"
	"net/http"
)

type UserPresenterImpl struct{}

func NewUserPresenter() usecase.UserPresenter {
	return &UserPresenterImpl{}
}

func (p *UserPresenterImpl) ViewAll(ctx context.Context, output *output.UserGetAllOutputData, err error) {
	w, _ := ctx.Value("resWriter").(http.ResponseWriter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	returnJSON(w, output)
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
