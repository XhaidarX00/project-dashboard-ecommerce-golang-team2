package productservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type ProductService interface {
	CreateProduct(productInput models.Product) error
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
func (p *productService) CreateProduct(productInput models.Product) error {
	panic("unimplemented")
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
