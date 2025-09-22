package shop

import (
	"edot/internal/domain"
	"edot/internal/repositories/shop"
	"edot/pkg/postgres"
	"fmt"
	"strings"
)

type ShopUsecase struct {
	Repo *shop.ShopRepositoryImpl
}

func NewShopUsecase() *ShopUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := shop.NewShopRepository(databaseHandler)
	return &ShopUsecase{
		Repo: repo,
	}
}

func (uc *ShopUsecase) GetAll() ([]domain.Shop, error) {
	return uc.Repo.List()
}

func (uc *ShopUsecase) GetByID(id int64) (*domain.Shop, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	shop, err := uc.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return shop, nil
}

func (uc *ShopUsecase) Create(s *domain.Shop) error {
	if s == nil {
		return fmt.Errorf("shop cannot be nil")
	}
	if err := uc.validateShop(s); err != nil {
		return err
	}
	return uc.Repo.Create(s)
}

func (uc *ShopUsecase) validateShop(s *domain.Shop) error {
	name := strings.TrimSpace(s.Name)
	if name == "" {
		return fmt.Errorf("shop name cannot be empty")
	}
	s.Name = name
	return nil
}
