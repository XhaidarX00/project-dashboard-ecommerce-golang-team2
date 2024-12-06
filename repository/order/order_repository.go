package orderrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository interface {
	UpdateStatus(id int, status string) error
	GetByID(id int) (*models.Order, error)
	GetAll() ([]models.Order, error)
	CountOrder() (int, error)
	CountTotalPriceOrder() (int, error)
	GetEarningEachMonth() (int, error)
}

type orderRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// GetEarningEachMonth implements OrderRepository.
func (o *orderRepository) GetEarningEachMonth() (int, error) {
	panic("unimplemented")
}

// CountTotalPriceOrder implements OrderRepository.
func (o *orderRepository) CountTotalPriceOrder() (int, error) {
	panic("unimplemented")
}

// CountOrder implements OrderRepository.
func (o *orderRepository) CountOrder() (int, error) {
	panic("unimplemented")
}

// GetAll implements OrderRepository.
func (o *orderRepository) GetAll() ([]models.Order, error) {
	panic("unimplemented")
}

// GetByID implements OrderRepository.
func (o *orderRepository) GetByID(id int) (*models.Order, error) {
	panic("unimplemented")
}

// UpdateStatus implements OrderRepository.
func (o *orderRepository) UpdateStatus(id int, status string) error {
	panic("unimplemented")
}

func NewOrderRepository(db *gorm.DB, log *zap.Logger) OrderRepository {
	return &orderRepository{DB: db, Log: log}
}
