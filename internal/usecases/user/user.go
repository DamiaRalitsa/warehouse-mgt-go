package user

import (
	"edot/internal/domain"
	"edot/internal/repositories/user"
	"edot/pkg/postgres"
	"fmt"
	"strings"
)

type UserUsecase struct {
	Repo *user.UserRepositoryImpl
}

func NewUserUsecase() *UserUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := user.NewUserRepository(databaseHandler)
	return &UserUsecase{
		Repo: repo,
	}
}

func (uc *UserUsecase) GetAll() ([]domain.User, error) {
	return uc.Repo.List()
}

func (uc *UserUsecase) GetByID(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}
	user, err := uc.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) Create(u *domain.User) error {
	if u == nil {
		return fmt.Errorf("user cannot be nil")
	}
	if err := uc.validateUser(u); err != nil {
		return err
	}
	return uc.Repo.Create(u)
}

func (uc *UserUsecase) validateUser(u *domain.User) error {
	name := strings.TrimSpace(u.Name)
	if name == "" {
		return fmt.Errorf("user name cannot be empty")
	}
	u.Name = name
	// Validate phone or email
	if strings.TrimSpace(u.Phone) == "" && strings.TrimSpace(u.Email) == "" {
		return fmt.Errorf("either phone or email must be provided")
	}
	return nil
}

// Login authenticates a user by phone or email
func (uc *UserUsecase) Login(identifier string) (*domain.User, error) {
	if strings.TrimSpace(identifier) == "" {
		return nil, fmt.Errorf("identifier (phone/email) required")
	}
	// Try phone first
	user, err := uc.Repo.GetByPhoneOrEmail(identifier)
	if err != nil {
		return nil, err
	}
	return user, nil
}
