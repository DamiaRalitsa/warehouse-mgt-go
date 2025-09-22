package shop

import "edot/internal/domain"

type ShopUsecase struct {
	Repo domain.ShopRepository
}

func (uc *ShopUsecase) Create(shop *domain.Shop) error {
	return uc.Repo.Create(shop)
}

func (uc *ShopUsecase) GetByID(id int64) (*domain.Shop, error) {
	return uc.Repo.GetByID(id)
}

func (uc *ShopUsecase) List() ([]domain.Shop, error) {
	return uc.Repo.List()
}
