package shop

import (
	"edot/internal/domain"
	"edot/pkg/postgres"
	"fmt"
	"log"
)

type ShopRepositoryImpl struct {
	db postgres.DatabaseHandlerFunc
}

func NewShopRepository(db postgres.DatabaseHandlerFunc) *ShopRepositoryImpl {
	return &ShopRepositoryImpl{db: db}
}

func (r *ShopRepositoryImpl) Create(shop *domain.Shop) error {
	if shop == nil {
		return fmt.Errorf("shop cannot be nil")
	}
	query := `INSERT INTO shops (id, name) VALUES ($1, $2)`
	err := r.db(nil, true, query, shop.ID, shop.Name)
	if err != nil {
		log.Printf("Error inserting shop: %v\n", err)
		return err
	}
	return nil
}

func (r *ShopRepositoryImpl) GetByID(id int64) (*domain.Shop, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	var results []domain.Shop
	query := `SELECT id, name FROM shops WHERE id = $1`
	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving shop by ID: %v\n", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no shop found with id: %d", id)
	}
	return &results[0], nil
}

func (r *ShopRepositoryImpl) List() ([]domain.Shop, error) {
	var results []domain.Shop
	query := `SELECT id, name FROM shops ORDER BY name ASC`
	err := r.db(&results, false, query)
	if err != nil {
		log.Printf("Error retrieving shops: %v\n", err)
		return []domain.Shop{}, err
	}
	if results == nil {
		results = []domain.Shop{}
	}
	return results, nil
}
