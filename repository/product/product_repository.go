package productrepository

import (
	"dashboard-ecommerce-team2/models"
	"encoding/json"
	"fmt"

	// "encoding/json"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(productInput *models.Product) error
	Update(productInput models.Product) error
	Delete(id int) error
	GetByID(id int) (*models.ProductID, error)
	GetAll(page, pageSize int) ([]*models.ProductWithCategory, int64, error)
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
	var product models.Product

	err := p.DB.First(&product, id).Error
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("product not found")
		}
		p.Log.Error("Failed to fetch product for deletion", zap.Error(err))
		return err
	}

	err = p.DB.Delete(&product).Error
	if err != nil {
		p.Log.Error("Failed to delete product", zap.Error(err))
		return err
	}

	return nil
}

// GetAll implements ProductRepository.
func (p *productRepository) GetAll(page, pageSize int) ([]*models.ProductWithCategory, int64, error) {
	// var products []*models.Product
	var productsWithCategory []*models.ProductWithCategory
	var totalItems int64

	offset := (page - 1) * pageSize

	// Menghitung total items
	err := p.DB.Model(&models.Product{}).Count(&totalItems).Error
	if err != nil {
		p.Log.Error("Failed to count total products", zap.Error(err))
		return nil, 0, err
	}

	// Mengambil produk dengan pagination
	err = p.DB.Table("products").
		Select(`products.id, 
            categories.name AS category_name, 
            products.name, 
            products.code_product, 
            products.images, 
            products.description, 
            products.stock, 
            products.price, 
            products.published`).
		Joins("JOIN categories ON categories.id = products.category_id").
		Offset(offset).
		Limit(pageSize).
		Find(&productsWithCategory).Error

	if err != nil {
		p.Log.Error("Failed to fetch products", zap.Error(err))
		return nil, 0, err
	}

	return productsWithCategory, totalItems, nil
}

// GetByID implements ProductRepository.
func (p *productRepository) GetByID(id int) (*models.ProductID, error) {
	var productID models.ProductID
	var variantJSON []byte

	err := p.DB.Table("products").
		Select(`
			products.id, 
			categories.name AS category_name, 
			CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant,
			products.name, 
			products.code_product, 
			products.images, 
			products.description, 
			products.stock, 
			products.price, 
			products.published
		`).
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.id = ?", id).
		First(&productID).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("product not found")
		}
		p.Log.Error("Failed to fetch product by ID", zap.Error(err))
		return nil, err
	}
	if len(variantJSON) > 0 {
		if err := json.Unmarshal(variantJSON, &productID.Variant); err != nil {
			p.Log.Error("Repository: failed to unmarshal variant JSON", zap.Error(err))
			return nil, err
		}
	}

	return &productID, nil
}

// Update implements ProductRepository.
func (p *productRepository) Update(productInput models.Product) error {
	panic("unimplemented")
}
