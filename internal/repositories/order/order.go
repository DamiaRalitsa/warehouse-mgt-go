package order

import (
	"edot/internal/domain"
	"edot/pkg/postgres"
	"fmt"
	"log"
	"time"
)

type OrderRepositoryImpl struct {
	db postgres.DatabaseHandlerFunc
}

func NewOrderRepository(db postgres.DatabaseHandlerFunc) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Create(order *domain.Order) error {
	if order == nil {
		return fmt.Errorf("order cannot be nil")
	}
	query := `INSERT INTO orders (id, user_id, product_id, shop_id) VALUES ($1, $2, $3, $4)`
	err := r.db(nil, true, query, order.ID, order.UserID, order.ProductID, order.ShopID)
	if err != nil {
		log.Printf("Error inserting order: %v\n", err)
		return err
	}
	return nil
}

func (r *OrderRepositoryImpl) GetByID(id int64) (*domain.Order, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	var results []domain.Order
	query := `SELECT id, user_id, product_id, shop_id FROM orders WHERE id = $1`
	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving order by ID: %v\n", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no order found with id: %d", id)
	}
	return &results[0], nil
}

func (r *OrderRepositoryImpl) List() ([]domain.Order, error) {
	var results []domain.Order
	query := `SELECT id, user_id, product_id, shop_id FROM orders ORDER BY id ASC`
	err := r.db(&results, false, query)
	if err != nil {
		log.Printf("Error retrieving orders: %v\n", err)
		return []domain.Order{}, err
	}
	if results == nil {
		results = []domain.Order{}
	}
	return results, nil
}

func (r *OrderRepositoryImpl) ReserveStock(productID int64, quantity int) error {
	query := `UPDATE products SET reserved = reserved + $1, stock = stock - $1 WHERE id = $2 AND stock >= $1`
	err := r.db(nil, true, query, quantity, productID)
	if err != nil {
		log.Printf("Error reserving stock: %v\n", err)
		return err
	}
	return nil
}

func (r *OrderRepositoryImpl) ReleaseStock(productID int64, quantity int) error {
	query := `UPDATE products SET reserved = reserved - $1, stock = stock + $1 WHERE id = $2 AND reserved >= $1`
	err := r.db(nil, true, query, quantity, productID)
	if err != nil {
		log.Printf("Error releasing stock: %v\n", err)
		return err
	}
	return nil
}

func (r *OrderRepositoryImpl) ListExpiredUnpaidOrders(expirySeconds int64) ([]domain.Order, error) {
	var results []domain.Order
	now := time.Now().Unix()
	query := `SELECT id, user_id, product_id, shop_id, reserved_at, paid FROM orders WHERE paid = false AND reserved_at > 0 AND reserved_at < $1`
	err := r.db(&results, false, query, now-expirySeconds)
	if err != nil {
		log.Printf("Error retrieving expired unpaid orders: %v\n", err)
		return []domain.Order{}, err
	}
	if results == nil {
		results = []domain.Order{}
	}
	return results, nil
}
