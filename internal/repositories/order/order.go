package order

import "edot/internal/domain"

type OrderRepositoryImpl struct{}

func (r *OrderRepositoryImpl) Create(order *domain.Order) error {
	return nil
}

func (r *OrderRepositoryImpl) GetByID(id int64) (*domain.Order, error) {
	return &domain.Order{ID: id, UserID: 1, ProductID: 1, ShopID: 1}, nil
}

func (r *OrderRepositoryImpl) List() ([]domain.Order, error) {
	return []domain.Order{}, nil
}
