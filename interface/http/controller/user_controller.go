package controller

import (
	"haiken/usecase"
	"net/http"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	ui usecase.UserInteractor
}

func NewUserController(ui usecase.UserInteractor) UserController {
	return &UserControllerImpl{
		ui: ui,
	}
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *UserControllerImpl) GetByID(w http.ResponseWriter, r *http.Request) {

}