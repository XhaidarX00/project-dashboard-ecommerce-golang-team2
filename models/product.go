package models

import (
	// "database/sql/driver"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	// "errors"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int             `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	CategoryID  int             `json:"category_id" gorm:"not null" binding:"required,gt=0" example:"1"`
	Name        string          `json:"name" gorm:"size:100;not null" binding:"required,min=3" example:"Smartphone"`
	CodeProduct string          `json:"code_product" gorm:"size:100;unique;not null" binding:"required" example:"SPH-001"`
	Images      StringArray     `json:"images" form:"file" gorm:"type:jsonb;default:'[]'" binding:"omitempty,dive,url" example:"[\"/images/smartphone1.png\"]"`
	Description string          `json:"description" binding:"omitempty" example:"Latest smartphone with advanced features"`
	Stock       int             `json:"stock" gorm:"not null" binding:"required,gt=0" example:"50"`
	Price       float64         `json:"price" gorm:"type:decimal(10,2);not null" binding:"required,gt=0" example:"699.99"`
	Published   bool            `json:"published" gorm:"type:boolean;default:false" example:"true"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

// type ImageJSON []string
type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}
	byteValue, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringArray")
	}
	return json.Unmarshal(byteValue, s)
}

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (v *Product) BeforeSave(tx *gorm.DB) (err error) {
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()

	return nil
}

type StockHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Type      string    `gorm:"size:10;check:type IN ('in','out')" json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type BestProduct struct {
	Name  string `json:"name"`
	Total string `json:"total"`
}

func ProductSeed() []Product {
	return []Product{
		{
			CategoryID:  1,
			Name:        "Smartphone",
			CodeProduct: "SPH-001",
			Images:      []string{"/images/smartphone1.png", "/images/smartphone2.png"},
			Description: "Latest smartphone with advanced features",
			Stock:       50,
			Price:       699.99,
			Published:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			CategoryID:  2,
			Name:        "Laptop",
			CodeProduct: "LAP-002",
			Images:      []string{"/images/laptop1.png", "/images/laptop2.png"},
			Description: "Powerful laptop with high performance",
			Stock:       30,
			Price:       1299.99,
			Published:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			CategoryID:  1,
			Name:        "Smartwatch",
			CodeProduct: "SW-003",
			Images:      []string{"/images/smartwatch1.png", "/images/smartwatch2.png"},
			Description: "Smartwatch with health monitoring features",
			Stock:       100,
			Price:       199.99,
			Published:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			CategoryID:  3,
			Name:        "Headphones",
			CodeProduct: "HP-004",
			Images:      []string{"/images/headphones1.png", "/images/headphones2.png"},
			Description: "Wireless headphones with noise cancellation",
			Stock:       150,
			Price:       89.99,
			Published:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
}

func StockHistorySeed() []StockHistory {
	return []StockHistory{}
}
