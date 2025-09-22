package user

import (
	"edot/internal/domain"
	"edot/pkg/postgres"
	"fmt"
	"log"
)

type UserRepositoryImpl struct {
	db postgres.DatabaseHandlerFunc
}

func NewUserRepository(db postgres.DatabaseHandlerFunc) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}

	query := `INSERT INTO users (id, name) VALUES ($1, $2)`
	err := r.db(nil, true, query, user.ID, user.Name)
	if err != nil {
		log.Printf("Error inserting user: %v\n", err)
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) GetByID(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	var results []domain.User
	query := `SELECT id, name FROM users WHERE id = $1`
	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving user by ID: %v\n", err)
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no user found with id: %d", id)
	}
	return &results[0], nil
}

func (r *UserRepositoryImpl) List() ([]domain.User, error) {
	var results []domain.User
	query := `SELECT id, name FROM users ORDER BY name ASC`
	err := r.db(&results, false, query)
	if err != nil {
		log.Printf("Error retrieving users: %v\n", err)
		return []domain.User{}, err
	}
	if results == nil {
		results = []domain.User{}
	}
	return results, nil
}
