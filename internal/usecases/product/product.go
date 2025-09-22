package product

import "edot/internal/domain"

type ProductUsecase struct {
	Repo domain.ProductRepository
}

func (uc *ProductUsecase) Create(product *domain.Product) error {
	return uc.Repo.Create(product)
}

func (uc *ProductUsecase) GetByID(id int64) (*domain.Product, error) {
	return uc.Repo.GetByID(id)
}

func (uc *ProductUsecase) List() ([]domain.Product, error) {
	return uc.Repo.List()
}
