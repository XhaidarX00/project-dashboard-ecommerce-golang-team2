package stockrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StockRepository interface {
	Update(StockInput models.StockHistory) error
	Delete(id int) error
	GetByID(id int) (*models.StockHistory, error)
}

type stockRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Delete implements StockRepository.
func (s *stockRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetByID implements StockRepository.
func (s *stockRepository) GetByID(id int) (*models.StockHistory, error) {
	panic("unimplemented")
}

// UpdateStock implements StockRepository.
func (s *stockRepository) Update(StockInput models.StockHistory) error {
	panic("unimplemented")
}

func NewStockRepository(db *gorm.DB, log *zap.Logger) StockRepository {
	return &stockRepository{DB: db, Log: log}
}
