package product

import (
	"edot/internal/domain"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	uc := ProductUsecase{Repo: &mockProductRepo{}}
	product := &domain.Product{ID: 1, Name: "Test"}
	if err := uc.Create(product); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

type mockProductRepo struct{}

func (m *mockProductRepo) Create(product *domain.Product) error { return nil }
func (m *mockProductRepo) GetByID(id int64) (*domain.Product, error) {
	return &domain.Product{ID: id, Name: "Test"}, nil
}
func (m *mockProductRepo) List() ([]domain.Product, error) { return []domain.Product{}, nil }
