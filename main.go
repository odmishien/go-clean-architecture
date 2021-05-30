package main

import (
	"haiken/interface/http/controller"
	"haiken/registry"
	"log"
	"net/http"
)

func main() {
	ui := registry.NewUserInteractor()
	uc := controller.NewUserController(ui)

	http.HandleFunc("/get-users", uc.GetAll)
	http.HandleFunc("/user", uc.Create)
	http.HandleFunc("/get-user", uc.GetByID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
