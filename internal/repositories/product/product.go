package product

import (
	"edot/internal/domain"
	"edot/pkg/postgres"
	"fmt"
	"log"
)

type ProductRepositoryImpl struct {
	db postgres.DatabaseHandlerFunc
}

func NewProductRepository(db postgres.DatabaseHandlerFunc) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) Create(product *domain.Product) error {
	if product == nil {
		return fmt.Errorf("product cannot be nil")
	}
	query := `INSERT INTO products (id, name) VALUES ($1, $2)`
	err := r.db(nil, true, query, product.ID, product.Name)
	if err != nil {
		log.Printf("Error inserting product: %v\n", err)
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) GetByID(id int64) (*domain.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	var results []domain.Product
	query := `SELECT id, name FROM products WHERE id = $1`
	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving product by ID: %v\n", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no product found with id: %d", id)
	}
	return &results[0], nil
}

func (r *ProductRepositoryImpl) List() ([]domain.Product, error) {
	var results []domain.Product
	query := `SELECT id, name FROM products ORDER BY name ASC`
	err := r.db(&results, false, query)
	if err != nil {
		log.Printf("Error retrieving products: %v\n", err)
		return []domain.Product{}, err
	}
	if results == nil {
		results = []domain.Product{}
	}
	return results, nil
}
