package warehouse

import (
	"edot/internal/domain"
	"edot/internal/repositories/warehouse"
	"edot/pkg/postgres"
	"fmt"
	"strings"
)

type WarehouseUsecase struct {
	Repo domain.WarehouseRepository
}

func NewWarehouseUsecase() *WarehouseUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := warehouse.NewWarehouseRepository(databaseHandler)
	return &WarehouseUsecase{
		Repo: repo,
	}
}

func (uc *WarehouseUsecase) GetAll() ([]domain.Warehouse, error) {
	return uc.Repo.List()
}

func (uc *WarehouseUsecase) GetByID(id int64) (*domain.Warehouse, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	warehouse, err := uc.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (uc *WarehouseUsecase) Create(w *domain.Warehouse) error {
	if w == nil {
		return fmt.Errorf("warehouse cannot be nil")
	}
	if err := uc.validateWarehouse(w); err != nil {
		return err
	}
	return uc.Repo.Create(w)
}

func (uc *WarehouseUsecase) validateWarehouse(w *domain.Warehouse) error {
	name := strings.TrimSpace(w.Name)
	if name == "" {
		return fmt.Errorf("warehouse name cannot be empty")
	}
	w.Name = name
	return nil
}
