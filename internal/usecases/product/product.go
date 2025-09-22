package product

import (
	"edot/internal/domain"
	"edot/internal/repositories/product"
	"edot/pkg/postgres"
	"fmt"
	"strings"
)

type ProductUsecase struct {
	Repo *product.ProductRepositoryImpl
}

func NewProductUsecase() *ProductUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := product.NewProductRepository(databaseHandler)
	return &ProductUsecase{
		Repo: repo,
	}
}

func (uc *ProductUsecase) GetAll() ([]domain.Product, error) {
	return uc.Repo.List()
}

func (uc *ProductUsecase) GetByID(id int64) (*domain.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	product, err := uc.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (uc *ProductUsecase) Create(p *domain.Product) error {
	if p == nil {
		return fmt.Errorf("product cannot be nil")
	}
	if err := uc.validateProduct(p); err != nil {
		return err
	}
	return uc.Repo.Create(p)
}

func (uc *ProductUsecase) validateProduct(p *domain.Product) error {
	name := strings.TrimSpace(p.Name)
	if name == "" {
		return fmt.Errorf("product name cannot be empty")
	}
	p.Name = name
	return nil
}
