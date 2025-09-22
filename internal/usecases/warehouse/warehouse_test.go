package warehouse

import (
	"edot/internal/domain"
	"testing"
)

func TestCreateWarehouse(t *testing.T) {
	uc := WarehouseUsecase{Repo: &mockWarehouseRepo{}}
	warehouse := &domain.Warehouse{ID: 1, Name: "Test"}
	if err := uc.Create(warehouse); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

type mockWarehouseRepo struct{}

func (m *mockWarehouseRepo) Create(warehouse *domain.Warehouse) error { return nil }
func (m *mockWarehouseRepo) GetByID(id int64) (*domain.Warehouse, error) {
	return &domain.Warehouse{ID: id, Name: "Test"}, nil
}
func (m *mockWarehouseRepo) List() ([]domain.Warehouse, error) { return []domain.Warehouse{}, nil }
