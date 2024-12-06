package userrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userInput models.User) error
	GetByEmail(email string) (*models.User, error)
	UpdatePassword(id, newPassword string) error
	CountCustomer() (int, error)
}

type userRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements UserRepository.
func (u *userRepository) Create(userInput models.User) error {
	panic("unimplemented")
}

// GetByEmail implements UserRepository.
func (u *userRepository) GetByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// UpdatePassword implements UserRepository.
func (u *userRepository) UpdatePassword(id, newPassword string) error {
	panic("unimplemented")
}

// CountCustomer implements UserRepository.

func (u *userRepository) CountCustomer() (int, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{DB: db, Log: log}
}
