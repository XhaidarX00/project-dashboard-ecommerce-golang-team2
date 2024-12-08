package productrepository

import (
	"dashboard-ecommerce-team2/models"
	"time"

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
	CountEachProduct() ([]models.BestProduct, error)
}

type productRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewProductRepository(db *gorm.DB, log *zap.Logger) ProductRepository {
	return &productRepository{DB: db, Log: log}
}

// CountEachProduct implements ProductRepository.
func (p *productRepository) CountEachProduct() ([]models.BestProduct, error) {
	type ProductCount struct {
		ProductID uint   `gorm:"column:product_id"`
		Name      string `gorm:"column:name"`
		Total     int    `gorm:"column:total"`
	}

	var productCounts []ProductCount
	var bestProducts []models.BestProduct

	// Perform the query to count products grouped by product ID
	err := p.DB.Table("order_items").
		Select("order_items.product_id, products.name, COALESCE(COUNT(order_items.product_id), 0) AS total").
		Joins("JOIN orders ON orders.id = order_items.order_id AND orders.status = ?", "completed").
		Joins("JOIN products ON products.id = order_items.product_id").
		Group("order_items.product_id, products.name").
		Order("total DESC").
		Scan(&productCounts).Error

	if err != nil {
		return nil, err
	}

	// Map the query results to the []models.BestProduct slice
	for _, pc := range productCounts {
		bestProducts = append(bestProducts, models.BestProduct{
			ProductID: int(pc.ProductID),
			Name:      pc.Name,
			Total:     pc.Total,
		})
	}

	return bestProducts, nil
}

// CountProduct implements ProductRepository.
func (p *productRepository) CountProduct() (int, error) {
	var count int64

	// Query to count distinct products in completed orders
	err := p.DB.Table("order_items").
		Joins("JOIN orders ON orders.id = order_items.order_id AND orders.status = ?", "completed").
		Where("EXTRACT(MONTH FROM orders.created_at) = ? AND EXTRACT(YEAR FROM orders.created_at) = ?", time.Now().Month(), time.Now().Year()).
		Select("COALESCE(SUM(order_items.quantity),0)").
		Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
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
