package input

type UserCreateInputData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserGetByIDInputData struct {
	ID int
}
