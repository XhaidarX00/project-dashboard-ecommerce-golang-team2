package stockrepository

import (
	"dashboard-ecommerce-team2/models"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StockRepository interface {
	Update(StockInput models.StockHistory) error
	Delete(id int) error
	GetByID(id int) (*models.StockHistoryResponse, error)
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
func (s *stockRepository) GetByID(id int) (*models.StockHistoryResponse, error) {
	var stockHistory models.StockHistoryResponse
	var variantJSON []byte
	
	err := s.DB.Table("stock_histories").
		Select("stock_histories.id, stock_histories.product_id, stock_histories.type, stock_histories.created_at, stock_histories.updated_at, products.name as product_name, products.stock as product_stock, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant").
		Joins("JOIN products ON products.id = stock_histories.product_id").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("stock_histories.id = ?", id).
		First(&stockHistory).Error

		if err != nil {
			if err.Error() == "record not found" {
				return nil, fmt.Errorf("product not found")
			}
			s.Log.Error("Failed to fetch product by ID", zap.Error(err))
			return nil, err
		}
		if len(variantJSON) > 0 {
			if err := json.Unmarshal(variantJSON, &stockHistory.Variant); err != nil {
				s.Log.Error("Repository: failed to unmarshal variant JSON", zap.Error(err))
				return nil, err
			}
		}

	return &stockHistory, nil
}

// UpdateStock implements StockRepository.
func (s *stockRepository) Update(StockInput models.StockHistory) error {
	panic("unimplemented")
}

func NewStockRepository(db *gorm.DB, log *zap.Logger) StockRepository {
	return &stockRepository{DB: db, Log: log}
}
