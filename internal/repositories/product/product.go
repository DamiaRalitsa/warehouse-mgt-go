package product

import "edot/internal/domain"

type ProductRepositoryImpl struct{}

func (r *ProductRepositoryImpl) Create(product *domain.Product) error {
	return nil
}

func (r *ProductRepositoryImpl) GetByID(id int64) (*domain.Product, error) {
	return &domain.Product{ID: id, Name: "Sample Product"}, nil
}

func (r *ProductRepositoryImpl) List() ([]domain.Product, error) {
	return []domain.Product{}, nil
}
