package output

import "haiken/entity"

type UserGetAllOutputData struct {
	Users entity.Users `json:"users"`
}

type UserGetByIDOutputData struct {
	User entity.User `json:"user"`
}

type UserCreateOutputData struct {
	CreatedID int `json:"createdId"`
}
