package usecase

import "haiken/entity"

type UserInteractor interface {
	GetAll() (entity.Users, error)
	Create(id int, name, email, password string) error
	GetByID(id int) (*entity.User, error)
}

type UserInteractorImpl struct {
	repo UserRepository
}

func NewUserInteractor(repo UserRepository) UserInteractor {
	return &UserInteractorImpl{
		repo: repo,
	}
}

func (i *UserInteractorImpl) GetAll() (entity.Users, error) {
	return i.repo.GetAll()
}
func (i *UserInteractorImpl) Create(id int, name, email, password string) error {
	return nil
}

func (i *UserInteractorImpl) GetByID(id int) (*entity.User, error) {
	return nil, nil
}
