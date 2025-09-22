package warehouse

import "edot/internal/domain"

type WarehouseRepositoryImpl struct{}

func (r *WarehouseRepositoryImpl) Create(warehouse *domain.Warehouse) error {
	return nil
}

func (r *WarehouseRepositoryImpl) GetByID(id int64) (*domain.Warehouse, error) {
	return &domain.Warehouse{ID: id, Name: "Sample Warehouse"}, nil
}

func (r *WarehouseRepositoryImpl) List() ([]domain.Warehouse, error) {
	return []domain.Warehouse{}, nil
}
