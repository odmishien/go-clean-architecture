package usecase

type UserCreateInputData struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserGetByIDInputData struct {
	ID int
}
