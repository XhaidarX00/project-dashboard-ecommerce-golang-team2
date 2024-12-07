package productrepository

import (
	"dashboard-ecommerce-team2/models"
	// "encoding/json"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(productInput *models.Product) error
	Update(productInput models.Product) error
	Delete(id int) error
	GetByID(id int) (*models.Product, error)
	GetAll(page, pageSize int) ([]*models.Product, int64, error)
	CountProduct() (int, error)
	CountEachProduct() (int, error)
}

type productRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewProductRepository(db *gorm.DB, log *zap.Logger) ProductRepository {
	return &productRepository{DB: db, Log: log}
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
func (p *productRepository) Create(productInput *models.Product) error {
	
	p.Log.Info("Creating product", zap.Any("input", productInput))
	err := p.DB.Create(productInput).Error
	if err != nil {
		p.Log.Error("repository: Error creating product", zap.Error(err))
		return err
	}
	return nil
}

// Delete implements ProductRepository.
func (p *productRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements ProductRepository.
func (p *productRepository) GetAll(page, pageSize int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var totalItems int64

	offset := (page - 1) * pageSize

	// Menghitung total items
	err := p.DB.Model(&models.Product{}).Count(&totalItems).Error
	if err != nil {
		p.Log.Error("Failed to count total products", zap.Error(err))
		return nil, 0, err
	}

	// Mengambil produk dengan pagination
	err = p.DB.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		p.Log.Error("Failed to fetch products", zap.Error(err))
		return nil, 0, err
	}

	return products, totalItems, nil
}

// GetByID implements ProductRepository.
func (p *productRepository) GetByID(id int) (*models.Product, error) {
	panic("unimplemented")
}

// Update implements ProductRepository.
func (p *productRepository) Update(productInput models.Product) error {
	panic("unimplemented")
}


