package shop

import (
	"edot/internal/domain"
	"testing"
)

func TestCreateShop(t *testing.T) {
	uc := ShopUsecase{Repo: &mockShopRepo{}}
	shop := &domain.Shop{ID: 1, Name: "Test"}
	if err := uc.Create(shop); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

type mockShopRepo struct{}

func (m *mockShopRepo) Create(shop *domain.Shop) error { return nil }
func (m *mockShopRepo) GetByID(id int64) (*domain.Shop, error) {
	return &domain.Shop{ID: id, Name: "Test"}, nil
}
func (m *mockShopRepo) List() ([]domain.Shop, error) { return []domain.Shop{}, nil }
