package user

import (
	"edot/internal/domain"
	"testing"
)

func TestCreateUser(t *testing.T) {
	uc := UserUsecase{Repo: &mockUserRepo{}}
	user := &domain.User{ID: 1, Name: "Test"}
	if err := uc.Create(user); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

type mockUserRepo struct{}

func (m *mockUserRepo) Create(user *domain.User) error { return nil }
func (m *mockUserRepo) GetByID(id int64) (*domain.User, error) {
	return &domain.User{ID: id, Name: "Test"}, nil
}
func (m *mockUserRepo) List() ([]domain.User, error) { return []domain.User{}, nil }
