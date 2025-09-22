package warehouse

import (
	"edot/internal/domain"
	"edot/pkg/postgres"
	"fmt"
	"log"
)

type WarehouseRepositoryImpl struct {
	db postgres.DatabaseHandlerFunc
}

func NewWarehouseRepository(db postgres.DatabaseHandlerFunc) *WarehouseRepositoryImpl {
	return &WarehouseRepositoryImpl{db: db}
}

func (r *WarehouseRepositoryImpl) Create(warehouse *domain.Warehouse) error {
	if warehouse == nil {
		return fmt.Errorf("warehouse cannot be nil")
	}
	query := `INSERT INTO warehouses (id, name) VALUES ($1, $2)`
	err := r.db(nil, true, query, warehouse.ID, warehouse.Name)
	if err != nil {
		log.Printf("Error inserting warehouse: %v\n", err)
		return err
	}
	return nil
}

func (r *WarehouseRepositoryImpl) GetByID(id int64) (*domain.Warehouse, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	var results []domain.Warehouse
	query := `SELECT id, name FROM warehouses WHERE id = $1`
	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving warehouse by ID: %v\n", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no warehouse found with id: %d", id)
	}
	return &results[0], nil
}

func (r *WarehouseRepositoryImpl) List() ([]domain.Warehouse, error) {
	var results []domain.Warehouse
	query := `SELECT id, name FROM warehouses ORDER BY name ASC`
	err := r.db(&results, false, query)
	if err != nil {
		log.Printf("Error retrieving warehouses: %v\n", err)
		return []domain.Warehouse{}, err
	}
	if results == nil {
		results = []domain.Warehouse{}
	}
	return results, nil
}

func (r *WarehouseRepositoryImpl) ListByShopID(shopID int64) ([]domain.Warehouse, error) {
	var results []domain.Warehouse
	query := `SELECT id, name, shop_id FROM warehouses WHERE shop_id = $1 ORDER BY name ASC`
	err := r.db(&results, false, query, shopID)
	if err != nil {
		log.Printf("Error retrieving warehouses for shop: %v\n", err)
		return []domain.Warehouse{}, err
	}
	if results == nil {
		results = []domain.Warehouse{}
	}
	return results, nil
}

func (r *WarehouseRepositoryImpl) TransferProduct(productID int64, fromWarehouseID int64, toWarehouseID int64, quantity int) error {
	// Deduct from source warehouse
	deductQuery := `UPDATE warehouse_products SET quantity = quantity - $1 WHERE warehouse_id = $2 AND product_id = $3 AND quantity >= $1`
	err := r.db(nil, true, deductQuery, quantity, fromWarehouseID, productID)
	if err != nil {
		log.Printf("Error deducting product from warehouse: %v\n", err)
		return err
	}
	// Add to destination warehouse
	addQuery := `UPDATE warehouse_products SET quantity = quantity + $1 WHERE warehouse_id = $2 AND product_id = $3`
	err = r.db(nil, true, addQuery, quantity, toWarehouseID, productID)
	if err != nil {
		log.Printf("Error adding product to warehouse: %v\n", err)
		return err
	}
	return nil
}

func (r *WarehouseRepositoryImpl) SetActiveStatus(warehouseID int64, active bool) error {
	query := `UPDATE warehouses SET active = $1 WHERE id = $2`
	err := r.db(nil, true, query, active, warehouseID)
	if err != nil {
		log.Printf("Error updating warehouse status: %v\n", err)
		return err
	}
	return nil
}
