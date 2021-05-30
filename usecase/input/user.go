package input

type UserCreateInputData struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10"`
}

type UserGetByIDInputData struct {
	ID int
}
