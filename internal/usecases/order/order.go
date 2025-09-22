package order

import (
	"edot/internal/domain"
	"edot/internal/repositories/order"
	"edot/pkg/postgres"
	"fmt"
)

type OrderUsecase struct {
	Repo *order.OrderRepositoryImpl
}

func NewOrderUsecase() *OrderUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := order.NewOrderRepository(databaseHandler)
	return &OrderUsecase{
		Repo: repo,
	}
}

func (uc *OrderUsecase) GetAll() ([]domain.Order, error) {
	return uc.Repo.List()
}

func (uc *OrderUsecase) GetByID(id int64) (*domain.Order, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	order, err := uc.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (uc *OrderUsecase) Create(o *domain.Order) error {
	if o == nil {
		return fmt.Errorf("order cannot be nil")
	}
	if err := uc.validateOrder(o); err != nil {
		return err
	}
	return uc.Repo.Create(o)
}

func (uc *OrderUsecase) validateOrder(o *domain.Order) error {
	if o.UserID <= 0 {
		return fmt.Errorf("user_id must be greater than 0")
	}
	if o.ProductID <= 0 {
		return fmt.Errorf("product_id must be greater than 0")
	}
	if o.ShopID <= 0 {
		return fmt.Errorf("shop_id must be greater than 0")
	}
	return nil
}
