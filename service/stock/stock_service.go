package stockservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type StockService interface {
	UpdateProductStock(newStock int) error
	GetProductStockDetail() (*models.StockHistory, error)
	DeleteProductStock(id int) error
}

type stockService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// DeleteProductStock implements StockService.
func (s *stockService) DeleteProductStock(id int) error {
	panic("unimplemented")
}

// GetProductStockDetail implements StockService.
func (s *stockService) GetProductStockDetail() (*models.StockHistory, error) {
	panic("unimplemented")
}

// UpdateProductStock implements StockService.
func (s *stockService) UpdateProductStock(newStock int) error {
	panic("unimplemented")
}

func NewStockService(repo repository.Repository, log *zap.Logger) StockService {
	return &stockService{Repo: repo, Log: log}
}
