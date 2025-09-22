package order

import (
	"edot/internal/domain"
	orderrepo "edot/internal/repositories/order"
	"edot/pkg/postgres"
	"fmt"
	"time"
)

type OrderUsecase struct {
	Repo *orderrepo.OrderRepositoryImpl
}

func NewOrderUsecase() *OrderUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := orderrepo.NewOrderRepository(databaseHandler)
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

// Checkout: reserve stock for an order
func (u *OrderUsecase) Checkout(productID int64, quantity int) error {
	// Validate quantity
	if quantity <= 0 {
		return fmt.Errorf("invalid quantity")
	}
	// Reserve stock in repository
	err := u.Repo.ReserveStock(productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

// ReleaseStock: release reserved stock if payment not made
func (u *OrderUsecase) ReleaseStock(productID int64, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("invalid quantity")
	}
	err := u.Repo.ReleaseStock(productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (uc *OrderUsecase) ReleaseExpiredUnpaidOrders(expirySeconds int64) error {
	orders, err := uc.Repo.ListExpiredUnpaidOrders(expirySeconds)
	if err != nil {
		return err
	}
	for _, o := range orders {
		// Release stock for each expired order
		err := uc.ReleaseStock(o.ProductID, 1) // TODO: use actual quantity if tracked
		if err != nil {
			return err
		}
		// Optionally, mark order as released/cancelled in DB
	}
	return nil
}

func StartOrderReleaseBackgroundJob(uc *OrderUsecase, expirySeconds int64, intervalSeconds int64) {
	go func() {
		ticker := time.NewTicker(time.Duration(intervalSeconds) * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			uc.ReleaseExpiredUnpaidOrders(expirySeconds)
		}
	}()
}
