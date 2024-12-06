package productrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(productInput models.Product) error
	Update(productInput models.Product) error
	Delete(id int) error
	GetByID(id int) (*models.Product, error)
	GetAll() ([]models.Product, error)
	CountProduct() (int, error)
	CountEachProduct() (int, error)
}

type productRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// CountEachProduct implements ProductRepository.
func (p *productRepository) CountEachProduct() (int, error) {
	panic("unimplemented")
}

// CountProduct implements ProductRepository.
func (p *productRepository) CountProduct() (int, error) {
	panic("unimplemented")
}

// Create implements ProductRepository.
func (p *productRepository) Create(productInput models.Product) error {
	panic("unimplemented")
}

// Delete implements ProductRepository.
func (p *productRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements ProductRepository.
func (p *productRepository) GetAll() ([]models.Product, error) {
	panic("unimplemented")
}

// GetByID implements ProductRepository.
func (p *productRepository) GetByID(id int) (*models.Product, error) {
	panic("unimplemented")
}

// Update implements ProductRepository.
func (p *productRepository) Update(productInput models.Product) error {
	panic("unimplemented")
}

func NewProductRepository(db *gorm.DB, log *zap.Logger) ProductRepository {
	return &productRepository{DB: db, Log: log}
}
