package warehouse

import "edot/internal/domain"

type WarehouseUsecase struct {
	Repo domain.WarehouseRepository
}

func (uc *WarehouseUsecase) Create(warehouse *domain.Warehouse) error {
	return uc.Repo.Create(warehouse)
}

func (uc *WarehouseUsecase) GetByID(id int64) (*domain.Warehouse, error) {
	return uc.Repo.GetByID(id)
}

func (uc *WarehouseUsecase) List() ([]domain.Warehouse, error) {
	return uc.Repo.List()
}
