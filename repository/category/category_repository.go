package categoryrepository

import (
	"dashboard-ecommerce-team2/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(categoryInput models.Category) error
	Update(categoryInput models.Category) error
	Delete(id int) error
	GetByID(id int) (*models.Category, error)
	GetAll() ([]models.Category, error)
}

type categoryRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements CategoryRepository.
func (c *categoryRepository) Create(categoryInput models.Category) error {
	panic("unimplemented")
}

// Delete implements CategoryRepository.
func (c *categoryRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements CategoryRepository.
func (c *categoryRepository) GetAll() ([]models.Category, error) {
	panic("unimplemented")
}

// GetByID implements CategoryRepository.
func (c *categoryRepository) GetByID(id int) (*models.Category, error) {
	panic("unimplemented")
}

// Update implements CategoryRepository.
func (c *categoryRepository) Update(categoryInput models.Category) error {
	panic("unimplemented")
}

func NewCategoryRepository(db *gorm.DB, log *zap.Logger) CategoryRepository {
	return &categoryRepository{DB: db, Log: log}
}
