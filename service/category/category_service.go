package categoryservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type CategoryService interface {
	CreateCatergory(categoryInput models.Category) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	UpdateCategory(categoryInput models.Category) error
	DeleteCategory(id int) error
}

type categoryService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CreateCatergory implements CategoryService.
func (c *categoryService) CreateCatergory(categoryInput models.Category) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryService.
func (c *categoryService) DeleteCategory(id int) error {
	panic("unimplemented")
}

// GetAllCategories implements CategoryService.
func (c *categoryService) GetAllCategories() ([]models.Category, error) {
	panic("unimplemented")
}

// GetCategoryByID implements CategoryService.
func (c *categoryService) GetCategoryByID(id int) (*models.Category, error) {
	panic("unimplemented")
}

// UpdateCategory implements CategoryService.
func (c *categoryService) UpdateCategory(categoryInput models.Category) error {
	panic("unimplemented")
}

func NewCategoryService(repo repository.Repository, log *zap.Logger) CategoryService {
	return &categoryService{Repo: repo, Log: log}
}
