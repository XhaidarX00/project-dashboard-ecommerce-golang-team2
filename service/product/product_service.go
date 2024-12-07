package productservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	utils "dashboard-ecommerce-team2/util"
	"os"

	"go.uber.org/zap"
)

type ProductService interface {
	CreateProduct(product *models.Product, filePath string) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	UpdateProduct(productInput models.Product) error
	DeleteProduct(id int) error
}

type productService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(product *models.Product, filePath string) (*models.Product, error) {
	cdnURL, err := utils.UploadToCDN(filePath)
	if err != nil {
		p.Log.Error("service: upload failed", zap.Error(err))
		return product, err
	}

	product.Images = []string{cdnURL}

	if err := p.Repo.Product.Create(product); err != nil {
		p.Log.Error("service: create failed", zap.Error(err))
		return product, err
	}

	os.Remove(filePath)

	return product, nil
}

// DeleteProduct implements ProductService.
func (p *productService) DeleteProduct(id int) error {
	panic("unimplemented")
}

// GetAllProducts implements ProductService.
func (p *productService) GetAllProducts() ([]models.Product, error) {
	panic("unimplemented")
}

// GetProductByID implements ProductService.
func (p *productService) GetProductByID(id int) (*models.Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements ProductService.
func (p *productService) UpdateProduct(productInput models.Product) error {
	panic("unimplemented")
}

func NewProductService(repo repository.Repository, log *zap.Logger) ProductService {
	return &productService{Repo: repo, Log: log}
}
