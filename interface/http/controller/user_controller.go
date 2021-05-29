package controller

import (
	"fmt"
	"haiken/usecase"
	"log"
	"net/http"
)

// TODO: impl presentor
type UserController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
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

func (c *UserControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	us, err := c.ui.GetAll()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(us)
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *UserControllerImpl) GetByID(w http.ResponseWriter, r *http.Request) {

}
