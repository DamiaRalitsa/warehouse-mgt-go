package user

import "edot/internal/domain"

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	// TODO: implement DB logic
	return nil
}

func (r *UserRepositoryImpl) GetByID(id int64) (*domain.User, error) {
	// TODO: implement DB logic
	return &domain.User{ID: id, Name: "Sample User"}, nil
}

func (r *UserRepositoryImpl) List() ([]domain.User, error) {
	// TODO: implement DB logic
	return []domain.User{}, nil
}
