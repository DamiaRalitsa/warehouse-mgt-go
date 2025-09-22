package user

import "edot/internal/domain"

type UserUsecase struct {
	Repo domain.UserRepository
}

func (uc *UserUsecase) Create(user *domain.User) error {
	return uc.Repo.Create(user)
}

func (uc *UserUsecase) GetByID(id int64) (*domain.User, error) {
	return uc.Repo.GetByID(id)
}

func (uc *UserUsecase) List() ([]domain.User, error) {
	return uc.Repo.List()
}
