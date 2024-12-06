package orderservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type OrderService interface {
	UpdateOrderStatus(id int, status string) error
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
}

type orderService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// GetAllOrders implements OrderService.
func (o *orderService) GetAllOrders() ([]models.Order, error) {
	panic("unimplemented")
}

// GetOrderByID implements OrderService.
func (o *orderService) GetOrderByID(id int) (*models.Order, error) {
	panic("unimplemented")
}

// UpdateOrderStatus implements OrderService.
func (o *orderService) UpdateOrderStatus(id int, status string) error {
	panic("unimplemented")
}

func NewOrderService(repo repository.Repository, log *zap.Logger) OrderService {
	return &orderService{Repo: repo, Log: log}
}
