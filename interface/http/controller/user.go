package controller

import (
	"context"
	"haiken/usecase"
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
	ctx := r.Context()
	ctx = context.WithValue(ctx, "resWriter", w)
	c.ui.GetAll(ctx)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {

}
