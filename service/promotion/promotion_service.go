package promotionservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type PromotionService interface {
	GetAllPromotions() ([]models.Promotion, error)
	CreatePromotion(promoInput *models.Promotion) error
	UpdatePromotion(promoInput *models.Promotion) error
	DeletePromotion(id int) error
}

type promotionService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CreatePromotion implements PromotionService.
func (p *promotionService) CreatePromotion(promoInput *models.Promotion) error {
	panic("unimplemented")
}

// DeletePromotion implements PromotionService.
func (p *promotionService) DeletePromotion(id int) error {
	panic("unimplemented")
}

// GetAllPromotions implements PromotionService.
func (p *promotionService) GetAllPromotions() ([]models.Promotion, error) {
	panic("unimplemented")
}

// UpdatePromotion implements PromotionService.
func (p *promotionService) UpdatePromotion(promoInput *models.Promotion) error {
	panic("unimplemented")
}

func NewPromotionService(repo repository.Repository, log *zap.Logger) PromotionService {
	return &promotionService{Repo: repo, Log: log}
}
