package order

import (
	"edot/internal/domain"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	uc := OrderUsecase{Repo: &mockOrderRepo{}}
	order := &domain.Order{ID: 1, UserID: 1, ProductID: 1, ShopID: 1}
	if err := uc.Create(order); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

type mockOrderRepo struct{}

func (m *mockOrderRepo) Create(order *domain.Order) error { return nil }
func (m *mockOrderRepo) GetByID(id int64) (*domain.Order, error) {
	return &domain.Order{ID: id, UserID: 1, ProductID: 1, ShopID: 1}, nil
}
func (m *mockOrderRepo) List() ([]domain.Order, error) { return []domain.Order{}, nil }
