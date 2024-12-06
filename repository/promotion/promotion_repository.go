package promotionrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	Create(promotionInput models.Promotion) error
	Update(promotionInput models.Promotion) error
	Delete(id int) error
	GetAll() ([]models.Promotion, error)
}

type promotionRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements PromotionRepository.
func (p *promotionRepository) Create(promotionInput models.Promotion) error {
	panic("unimplemented")
}

// Delete implements PromotionRepository.
func (p *promotionRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements PromotionRepository.
func (p *promotionRepository) GetAll() ([]models.Promotion, error) {
	panic("unimplemented")
}

// Update implements PromotionRepository.
func (p *promotionRepository) Update(promotionInput models.Promotion) error {
	panic("unimplemented")
}

func NewPromotionRepository(db *gorm.DB, log *zap.Logger) PromotionRepository {
	return &promotionRepository{DB: db, Log: log}
}
