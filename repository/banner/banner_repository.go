package bannerrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BannerRepository interface {
	Create(bannerInput models.Banner) error
	Update(bannerInput models.Banner) error
	Delete(id int) error
	GetByID(id int) (*models.Banner, error)
}

type bannerRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements BannerRepository.
func (b *bannerRepository) Create(bannerInput models.Banner) error {
	panic("unimplemented")
}

// Delete implements BannerRepository.
func (b *bannerRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetByID implements BannerRepository.
func (b *bannerRepository) GetByID(id int) (*models.Banner, error) {
	panic("unimplemented")
}

// Update implements BannerRepository.
func (b *bannerRepository) Update(bannerInput models.Banner) error {
	panic("unimplemented")
}

func NewBannerRepository(db *gorm.DB, log *zap.Logger) BannerRepository {
	return &bannerRepository{DB: db, Log: log}
}
