package controller

import (
	"encoding/json"
	"haiken/usecase"
	"haiken/usecase/input"
	"net/http"
)

type UserController struct {
	ui usecase.UserInteractor
}

func NewUserController(ui usecase.UserInteractor) UserController {
	return UserController{
		ui: ui,
	}
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := setResponseWriter(r.Context(), w)
	c.ui.GetAll(ctx)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	ctx := setResponseWriter(r.Context(), w)
	var ci input.UserCreateInputData
	err := json.NewDecoder(r.Body).Decode(&ci)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.ui.Create(ctx, ci)
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := setResponseWriter(r.Context(), w)
	var ci input.UserGetByIDInputData
	err := json.NewDecoder(r.Body).Decode(&ci)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.ui.GetByID(ctx, ci)
}
