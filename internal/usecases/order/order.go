package order

import "edot/internal/domain"

type OrderUsecase struct {
	Repo domain.OrderRepository
}

func (uc *OrderUsecase) Create(order *domain.Order) error {
	return uc.Repo.Create(order)
}

func (uc *OrderUsecase) GetByID(id int64) (*domain.Order, error) {
	return uc.Repo.GetByID(id)
}

func (uc *OrderUsecase) List() ([]domain.Order, error) {
	return uc.Repo.List()
}
