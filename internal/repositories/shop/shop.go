package shop

import "edot/internal/domain"

type ShopRepositoryImpl struct{}

func (r *ShopRepositoryImpl) Create(shop *domain.Shop) error {
	return nil
}

func (r *ShopRepositoryImpl) GetByID(id int64) (*domain.Shop, error) {
	return &domain.Shop{ID: id, Name: "Sample Shop"}, nil
}

func (r *ShopRepositoryImpl) List() ([]domain.Shop, error) {
	return []domain.Shop{}, nil
}
